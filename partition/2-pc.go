package partition

import (
	"context"

	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

func (ls *ListenServer) PrepareCommit(ctx context.Context, req *pbpartition.PrepareCommitRequest) (*pbpartition.PrepareCommitResponse, error) {
	if deleteMsg := req.GetDelete(); deleteMsg != nil {
		ls.lockedMessage = deleteMsg
	} else if setMsg := req.GetSet(); setMsg != nil {
		ls.lockedMessage = setMsg
	} else {
		return nil, ErrUnsupported2PCMsg
	}

	// lock a db until locked message is committed or aborted
	ls.isLocked = true
	return &pbpartition.PrepareCommitResponse{Ok: true}, nil
}

func (ls *ListenServer) AbortCommit(ctx context.Context, req *pbpartition.AbortCommitRequest) (*pbpartition.AbortCommitResponse, error) {
	ls.isLocked = false
	ls.lockedMessage = nil
	return &pbpartition.AbortCommitResponse{}, nil
}

func (ls *ListenServer) Commit(ctx context.Context, req *pbpartition.CommitRequest) (*pbpartition.CommitResponse, error) {
	if ls.lockedMessage == nil {
		return nil, ErrNoLockedMessage
	}

	if deleteMsg, ok := ls.lockedMessage.(*prototypes.DeleteRequest); ok {
		ls.Delete(ctx, deleteMsg)
	} else if setMsg, ok := ls.lockedMessage.(*prototypes.SetRequest); ok {
		ls.Set(ctx, setMsg)
	} else {
		return nil, ErrUnsupported2PCMsg
	}

	ls.isLocked = false
	ls.lockedMessage = nil
	return &pbpartition.CommitResponse{}, nil
}
