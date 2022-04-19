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

#include <string>

namespace milvus {
inline bool
PrefixMatch(const std::string& str, const std::string& prefix) {
    if (prefix.length() > str.length()) {
        return false;
    }
    for (int i = 0; i < prefix.length(); i++) {
        if (prefix[i] != str[i]) {
            return false;
        }
    }
    return true;
}

inline bool
PostfixMatch(const std::string& str, const std::string& postfix) {
    if (postfix.length() > str.length()) {
        return false;
    }
    int i = postfix.length() - 1;
    int j = str.length() - 1;
    for (; i >= 0; i--, j--) {
        if (postfix[i] != str[j]) {
            return false;
        }
    }
    return true;
}
}  // namespace milvus
