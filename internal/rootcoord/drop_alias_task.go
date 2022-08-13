package rootcoord

import (
	"context"

	"github.com/milvus-io/milvus/internal/proto/milvuspb"
)

type dropAliasTask struct {
	baseRedoTask
	Req *milvuspb.DropAliasRequest
}

func (t *dropAliasTask) Prepare(ctx context.Context) error {
	return nil
}

func (t *dropAliasTask) Execute(ctx context.Context) error {
	// drop alias is atomic enough.
	if err := t.core.ExpireMetaCache(ctx, []string{t.Req.GetAlias()}, InvalidCollectionID, t.GetTs()); err != nil {
		return err
	}
	return t.core.meta.DropAlias(ctx, t.Req.GetAlias(), t.GetTs())
}
