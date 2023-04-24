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

	"github.com/milvus-io/milvus/internal/util/typeutil"
)

type dropDatabaseTask struct {
	baseTask
}

func (t *dropDatabaseTask) Prepare(ctx context.Context) error {
	t.SetStep(typeutil.TaskStepPreExecute)
	return nil
}

func (t *dropDatabaseTask) Execute(ctx context.Context) error {
	t.SetStep(typeutil.TaskStepExecute)
	//if err := t.core.ExpireMetaCache(ctx, []string{t.Req.GetAlias(), t.Req.GetCollectionName()}, InvalidCollectionID, t.GetTs()); err != nil {
	//	return err
	//}
	return t.core.meta.DropDatabase(ctx, "", t.GetTs())
}
