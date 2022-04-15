// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License

#pragma once

#include <boost/algorithm/string/predicate.hpp>
#include <cstring>
#include <memory>
#include <random>
#include <knowhere/index/vector_index/VecIndex.h>
#include <knowhere/index/vector_index/adapter/VectorAdapter.h>
#include <knowhere/index/vector_index/VecIndexFactory.h>
#include <knowhere/index/vector_index/IndexIVF.h>

#include "Constants.h"
#include "common/Schema.h"
#include "query/SearchOnIndex.h"
#include "segcore/SegmentGrowingImpl.h"
#include "segcore/SegmentSealedImpl.h"
#include "segcore/Utils.h"

using boost::algorithm::starts_with;

namespace milvus::segcore {

struct GeneratedData {
  std::vector<idx_t> row_ids_;
  std::vector<Timestamp> timestamps_;
  InsertData* raw_;
  std::vector<FieldId> field_ids;
  SchemaPtr schema_;
  template <typename T>
  std::vector<T>
  get_col(FieldId field_id) const {
      std::vector<T> ret(raw_->num_rows());
      for (auto target_field_data : raw_->fields_data()) {
          if (field_id.get() != target_field_data.field_id()) {
              continue;
          }

          auto& field_meta = schema_->operator[](field_id);
          if (field_meta.is_vector()) {
              if (field_meta.get_data_type() == DataType::VECTOR_FLOAT) {
                  int len = raw_->num_rows() * field_meta.get_dim();
                  ret.resize(len);
                  auto src_data = reinterpret_cast<const T*>(target_field_data.vectors().float_vector().data().data());
                  std::copy_n(src_data, len, ret.data());
              } else if (field_meta.get_data_type() == DataType::VECTOR_BINARY) {
                  int len = raw_->num_rows() * (field_meta.get_dim() / 8);
                  ret.resize(len);
                  auto src_data = reinterpret_cast<const T*>(target_field_data.vectors().binary_vector().data());
                  std::copy_n(src_data, len, ret.data());
              } else {
                  PanicInfo("unsupported");
              }

              return std::move(ret);
          }
          switch (field_meta.get_data_type()) {
              case DataType::BOOL: {
                  auto src_data = reinterpret_cast<const T*>(target_field_data.scalars().bool_data().data().data());
                  std::copy_n(src_data, raw_->num_rows(), ret.data());
                  break;
              }
              case DataType::INT8:
              case DataType::INT16:
              case DataType::INT32: {
                  auto src_data = reinterpret_cast<const T*>(target_field_data.scalars().int_data().data().data());
                  std::copy_n(src_data, raw_->num_rows(), ret.data());
                  break;
              }
              case DataType::INT64: {
                  auto src_data = reinterpret_cast<const T*>(target_field_data.scalars().long_data().data().data());
                  std::copy_n(src_data, raw_->num_rows(), ret.data());
                  break;
              }
              case DataType::FLOAT: {
                  auto src_data = reinterpret_cast<const T*>(target_field_data.scalars().float_data().data().data());
                  std::copy_n(src_data, raw_->num_rows(), ret.data());
                  break;
              }
              case DataType::DOUBLE: {
                  auto src_data = reinterpret_cast<const T*>(target_field_data.scalars().double_data().data().data());
                  std::copy_n(src_data, raw_->num_rows(), ret.data());
                  break;
              }
              case DataType::VarChar: {
                  auto src_data = reinterpret_cast<const T*>(target_field_data.scalars().string_data().data().data());
                  std::copy_n(src_data, raw_->num_rows(), ret.data());
                  break;
              }
              default: {
                  PanicInfo("unsupported");
              }

          }
      }
      return std::move(ret);
  }

  std::unique_ptr<DataArray>
  get_col(FieldId field_id) const {
      for (auto target_field_data : raw_->fields_data()) {
          if (field_id.get() == target_field_data.field_id()) {
              return std::make_unique<DataArray>(target_field_data);
          }
      }

      PanicInfo("field id not find");
  }

private:
  GeneratedData() = default;
  friend GeneratedData
  DataGen(SchemaPtr schema, int64_t N, uint64_t seed, uint64_t ts_offset);
};

inline GeneratedData
DataGen(SchemaPtr schema, int64_t N, uint64_t seed = 42, uint64_t ts_offset = 0) {
    using std::vector;
    std::default_random_engine er(seed);
    std::normal_distribution<> distr(0, 1);
    int offset = 0;

    auto insert_data = std::make_unique<InsertData>();
    auto insert_cols = [&insert_data](auto& data, int64_t count, auto& field_meta) {
        auto array = milvus::segcore::CreateDataArrayFrom(data.data(), count, field_meta);
        insert_data->mutable_fields_data()->AddAllocated(array.release());
    };

    for (auto field_id : schema->get_field_ids()) {
        auto field_meta = schema->operator[](field_id);
        switch (field_meta.get_data_type()) {
            case engine::DataType::VECTOR_FLOAT: {
                auto dim = field_meta.get_dim();
                vector<float> final(dim * N);
                bool is_ip = starts_with(field_meta.get_name().get(), "normalized");
#pragma omp parallel for
                for (int n = 0; n < N; ++n) {
                    vector<float> data(dim);
                    float sum = 0;

                    std::default_random_engine er2(seed + n);
                    std::normal_distribution<> distr2(0, 1);
                    for (auto& x : data) {
                        x = distr2(er2) + offset;
                        sum += x * x;
                    }
                    if (is_ip) {
                        sum = sqrt(sum);
                        for (auto& x : data) {
                            x /= sum;
                        }
                    }

                    std::copy(data.begin(), data.end(), final.begin() + dim * n);
                }
                insert_cols(final, N, field_meta);
                break;
            }
            case engine::DataType::VECTOR_BINARY: {
                auto dim = field_meta.get_dim();
                Assert(dim % 8 == 0);
                vector<uint8_t> data(dim / 8 * N);
                for (auto& x : data) {
                    x = er();
                }
                insert_cols(data, N, field_meta);
                break;
            }
            case engine::DataType::INT64: {
                vector<int64_t> data(N);
                // begin with counter
                if (starts_with(field_meta.get_name().get(), "counter")) {
                    int64_t index = 0;
                    for (auto& x : data) {
                        x = index++;
                    }
                } else {
                    int i = 0;
                    for (auto& x : data) {
                        x = er() % (2 * N);
                        x = i;
                        i++;
                    }
                }
                insert_cols(data, N, field_meta);
                break;
            }
            case engine::DataType::INT32: {
                vector<int> data(N);
                for (auto& x : data) {
                    x = er() % (2 * N);
                }
                insert_cols(data, N, field_meta);
                break;
            }
            case engine::DataType::FLOAT: {
                vector<float> data(N);
                for (auto& x : data) {
                    x = distr(er);
                }
                insert_cols(data, N, field_meta);
                break;
            }
            case engine::DataType::DOUBLE: {
                vector<double> data(N);
                for (auto& x : data) {
                    x = distr(er);
                }
                insert_cols(data, N, field_meta);
                break;
            }
            case engine::DataType::VarChar: {
                vector<std::string> data(N);
                for (auto& x: data) {
                    x = std::to_string(er());
                }
                insert_cols(data, N, field_meta);
                break;
            }
            default: {
                throw std::runtime_error("unimplemented");
            }
        }
        ++offset;
    }

    GeneratedData res;
    res.schema_ = schema;
    res.raw_ = insert_data.release();
    res.raw_->set_num_rows(N);
    for (int i = 0; i < N; ++i) {
        res.row_ids_.push_back(i);
        res.timestamps_.push_back(i + ts_offset);
    }

    return std::move(res);
}

inline auto
CreatePlaceholderGroup(int64_t num_queries, int dim, int64_t seed = 42) {
    namespace ser = milvus::proto::milvus;
    ser::PlaceholderGroup raw_group;
    auto value = raw_group.add_placeholders();
    value->set_tag("$0");
    value->set_type(ser::PlaceholderType::FloatVector);
    std::normal_distribution<double> dis(0, 1);
    std::default_random_engine e(seed);
    for (int i = 0; i < num_queries; ++i) {
        std::vector<float> vec;
        for (int d = 0; d < dim; ++d) {
            vec.push_back(dis(e));
        }
        // std::string line((char*)vec.data(), (char*)vec.data() + vec.size() * sizeof(float));
        value->add_values(vec.data(), vec.size() * sizeof(float));
    }
    return raw_group;
}

inline auto
CreatePlaceholderGroupFromBlob(int64_t num_queries, int dim, const float* src) {
    namespace ser = milvus::proto::milvus;
    ser::PlaceholderGroup raw_group;
    auto value = raw_group.add_placeholders();
    value->set_tag("$0");
    value->set_type(ser::PlaceholderType::FloatVector);
    int64_t src_index = 0;

    for (int i = 0; i < num_queries; ++i) {
        std::vector<float> vec;
        for (int d = 0; d < dim; ++d) {
            vec.push_back(src[src_index++]);
        }
        // std::string line((char*)vec.data(), (char*)vec.data() + vec.size() * sizeof(float));
        value->add_values(vec.data(), vec.size() * sizeof(float));
    }
    return raw_group;
}

inline auto
CreateBinaryPlaceholderGroup(int64_t num_queries, int64_t dim, int64_t seed = 42) {
    assert(dim % 8 == 0);
    namespace ser = milvus::proto::milvus;
    ser::PlaceholderGroup raw_group;
    auto value = raw_group.add_placeholders();
    value->set_tag("$0");
    value->set_type(ser::PlaceholderType::BinaryVector);
    std::default_random_engine e(seed);
    for (int i = 0; i < num_queries; ++i) {
        std::vector<uint8_t> vec;
        for (int d = 0; d < dim / 8; ++d) {
            vec.push_back(e());
        }
        // std::string line((char*)vec.data(), (char*)vec.data() + vec.size() * sizeof(float));
        value->add_values(vec.data(), vec.size());
    }
    return raw_group;
}

inline auto
CreateBinaryPlaceholderGroupFromBlob(int64_t num_queries, int64_t dim, const uint8_t* ptr) {
    assert(dim % 8 == 0);
    namespace ser = milvus::proto::milvus;
    ser::PlaceholderGroup raw_group;
    auto value = raw_group.add_placeholders();
    value->set_tag("$0");
    value->set_type(ser::PlaceholderType::BinaryVector);
    for (int i = 0; i < num_queries; ++i) {
        std::vector<uint8_t> vec;
        for (int d = 0; d < dim / 8; ++d) {
            vec.push_back(*ptr);
            ++ptr;
        }
        // std::string line((char*)vec.data(), (char*)vec.data() + vec.size() * sizeof(float));
        value->add_values(vec.data(), vec.size());
    }
    return raw_group;
}

inline json
SearchResultToJson(const SearchResult& sr) {
    int64_t num_queries = sr.num_queries_;
    int64_t topk = sr.topk_;
    std::vector<std::vector<std::string>> results;
    for (int q = 0; q < num_queries; ++q) {
        std::vector<std::string> result;
        for (int k = 0; k < topk; ++k) {
            int index = q * topk + k;
            result.emplace_back(std::to_string(sr.seg_offsets_[index]) + "->" + std::to_string(sr.distances_[index]));
        }
        results.emplace_back(std::move(result));
    }
    return json{results};
};

inline void
SealedLoader(const GeneratedData& dataset, SegmentSealed& seg) {
    // TODO
    auto row_count = dataset.row_ids_.size();
    {
        LoadFieldDataInfo info;
        FieldMeta field_meta(FieldName("RowID"), RowFieldID, engine::DataType::INT64);
        auto array = CreateScalarDataArrayFrom(dataset.row_ids_.data(), row_count, field_meta);
        info.field_data = array.release();
        info.row_count = dataset.row_ids_.size();
        info.field_id = RowFieldID.get();  // field id for RowId
        seg.LoadFieldData(info);
    }
    {
        LoadFieldDataInfo info;
        FieldMeta field_meta(FieldName("Timestamp"), TimestampFieldID, engine::DataType::INT64);
        auto array = CreateScalarDataArrayFrom(dataset.timestamps_.data(), row_count, field_meta);
        info.field_data = array.release();
        info.row_count = dataset.timestamps_.size();
        info.field_id = TimestampFieldID.get();
        seg.LoadFieldData(info);
    }
    for (auto field_data : dataset.raw_->fields_data()) {
        LoadFieldDataInfo info;
        info.field_id = field_data.field_id();
        info.row_count = row_count;
        info.field_data = &field_data;
        seg.LoadFieldData(info);
    }
}

inline std::unique_ptr<SegmentSealed>
SealedCreator(SchemaPtr schema, const GeneratedData& dataset, const LoadIndexInfo& index_info) {
    auto segment = CreateSealedSegment(schema);
    SealedLoader(dataset, *segment);
    segment->LoadIndex(index_info);
    return segment;
}

inline knowhere::VecIndexPtr
GenIndexing(int64_t N, int64_t dim, const float* vec) {
    // {knowhere::IndexParams::nprobe, 10},
    auto conf = knowhere::Config{{knowhere::meta::DIM, dim},
                                 {knowhere::IndexParams::nlist, 1024},
                                 {knowhere::Metric::TYPE, knowhere::Metric::L2},
                                 {knowhere::meta::DEVICEID, 0}};
    auto database = knowhere::GenDataset(N, dim, vec);
    auto indexing = std::make_shared<knowhere::IVF>();
    indexing->Train(database, conf);
    indexing->AddWithoutIds(database, conf);
    return indexing;
}

}  // namespace milvus::segcore
