// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rootcoord

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/milvuspb"
	"github.com/milvus-io/milvus/internal/metastore/model"
	mockrootcoord "github.com/milvus-io/milvus/internal/rootcoord/mocks"
	"github.com/milvus-io/milvus/internal/util/funcutil"
)

func Test_createPartitionTask_Prepare(t *testing.T) {
	t.Run("invalid msg type", func(t *testing.T) {
		task := &createPartitionTask{
			baseTask: newBaseTask(context.TODO(), nil),
			Req:      &milvuspb.CreatePartitionRequest{Base: &commonpb.MsgBase{MsgType: commonpb.MsgType_DropCollection}},
		}
		err := task.Prepare(context.Background())
		assert.Error(t, err)
	})

	t.Run("failed to get collection meta", func(t *testing.T) {
		core := newTestCore(withInvalidMeta())
		task := &createPartitionTask{
			baseTask: newBaseTask(context.TODO(), core),
			Req:      &milvuspb.CreatePartitionRequest{Base: &commonpb.MsgBase{MsgType: commonpb.MsgType_CreatePartition}},
		}
		err := task.Prepare(context.Background())
		assert.Error(t, err)
	})

	t.Run("normal case", func(t *testing.T) {
		collectionName := funcutil.GenRandomStr()
		coll := &model.Collection{Name: collectionName}

		meta := mockrootcoord.NewIMetaTable(t)
		meta.On("GetCollectionByName",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(coll.Clone(), nil)

		core := newTestCore(withMeta(meta))
		task := &createPartitionTask{
			baseTask: newBaseTask(context.TODO(), core),
			Req:      &milvuspb.CreatePartitionRequest{Base: &commonpb.MsgBase{MsgType: commonpb.MsgType_CreatePartition}},
		}
		err := task.Prepare(context.Background())
		assert.NoError(t, err)
		assert.True(t, coll.Equal(*task.collMeta))
	})
}

func Test_createPartitionTask_Execute(t *testing.T) {
	t.Run("create duplicate partition", func(t *testing.T) {
		collectionName := funcutil.GenRandomStr()
		partitionName := funcutil.GenRandomStr()
		coll := &model.Collection{Name: collectionName, Partitions: []*model.Partition{{PartitionName: partitionName}}}
		task := &createPartitionTask{
			collMeta: coll,
			Req:      &milvuspb.CreatePartitionRequest{CollectionName: collectionName, PartitionName: partitionName},
		}
		err := task.Execute(context.Background())
		assert.NoError(t, err)
	})

	t.Run("create too many partitions", func(t *testing.T) {
		cfgMaxPartitionNum := Params.RootCoordCfg.MaxPartitionNum
		partitions := make([]*model.Partition, 0, cfgMaxPartitionNum)
		for i := int64(0); i < cfgMaxPartitionNum; i++ {
			partitions = append(partitions, &model.Partition{})
		}
		collectionName := funcutil.GenRandomStr()
		partitionName := funcutil.GenRandomStr()
		coll := &model.Collection{Name: collectionName, Partitions: partitions}
		task := &createPartitionTask{
			collMeta: coll,
			Req:      &milvuspb.CreatePartitionRequest{CollectionName: collectionName, PartitionName: partitionName},
		}
		err := task.Execute(context.Background())
		assert.Error(t, err)
	})

	t.Run("failed to allocate partition id", func(t *testing.T) {
		collectionName := funcutil.GenRandomStr()
		partitionName := funcutil.GenRandomStr()
		coll := &model.Collection{Name: collectionName, Partitions: []*model.Partition{}}
		core := newTestCore(withInvalidIDAllocator())
		task := &createPartitionTask{
			baseTask: newBaseTask(context.TODO(), core),
			collMeta: coll,
			Req:      &milvuspb.CreatePartitionRequest{CollectionName: collectionName, PartitionName: partitionName},
		}
		err := task.Execute(context.Background())
		assert.Error(t, err)
	})

	t.Run("failed to expire cache", func(t *testing.T) {
		collectionName := funcutil.GenRandomStr()
		partitionName := funcutil.GenRandomStr()
		coll := &model.Collection{Name: collectionName, Partitions: []*model.Partition{}}
		core := newTestCore(withValidIDAllocator(), withInvalidProxyManager())
		task := &createPartitionTask{
			baseTask: newBaseTask(context.TODO(), core),
			collMeta: coll,
			Req:      &milvuspb.CreatePartitionRequest{CollectionName: collectionName, PartitionName: partitionName},
		}
		err := task.Execute(context.Background())
		assert.Error(t, err)
	})

	t.Run("failed to add partition meta", func(t *testing.T) {
		collectionName := funcutil.GenRandomStr()
		partitionName := funcutil.GenRandomStr()
		coll := &model.Collection{Name: collectionName, Partitions: []*model.Partition{}}
		core := newTestCore(withValidIDAllocator(), withValidProxyManager(), withInvalidMeta())
		task := &createPartitionTask{
			baseTask: newBaseTask(context.TODO(), core),
			collMeta: coll,
			Req:      &milvuspb.CreatePartitionRequest{CollectionName: collectionName, PartitionName: partitionName},
		}
		err := task.Execute(context.Background())
		assert.Error(t, err)
	})

	t.Run("normal case", func(t *testing.T) {
		collectionName := funcutil.GenRandomStr()
		partitionName := funcutil.GenRandomStr()
		coll := &model.Collection{Name: collectionName, Partitions: []*model.Partition{}}

		meta := mockrootcoord.NewIMetaTable(t)
		meta.On("AddPartition",
			mock.Anything,
			mock.Anything,
		).Return(nil)

		core := newTestCore(withValidIDAllocator(), withValidProxyManager(), withMeta(meta))
		task := &createPartitionTask{
			baseTask: newBaseTask(context.TODO(), core),
			collMeta: coll,
			Req:      &milvuspb.CreatePartitionRequest{CollectionName: collectionName, PartitionName: partitionName},
		}
		err := task.Execute(context.Background())
		assert.NoError(t, err)
	})
}
