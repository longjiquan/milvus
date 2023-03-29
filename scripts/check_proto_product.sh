#!/usr/bin/env bash

# Licensed to the LF AI & Data foundation under one
# or more contributor license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership. The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SCRIPTS_DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"

cd ${SCRIPTS_DIR}

GO_SRC_DIR="${SCRIPTS_DIR}/$1"

if [[ $(uname -s) == "Darwin" ]]; then
  if ! brew --prefix --installed grep >/dev/null 2>&1; then
        brew install grep
  fi
  export PATH="/usr/local/opt/grep/libexec/gnubin:$PATH"
fi

check_result=$(git status | grep -E ".*pb.go|.*pb.cc|.*pb.h")
echo "check_result: $check_result"
if test -z "$check_result"; then
  exit 0
else
  echo "The go file or cpp file generated by proto are not latest!"
  exit 1
fi
