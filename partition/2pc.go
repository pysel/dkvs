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

	return &pbpartition.PrepareCommitResponse{Ok: true}, nil
}

func (ls *ListenServer) AbortCommit(ctx context.Context, req *pbpartition.AbortCommitRequest) (*pbpartition.AbortCommitResponse, error) {
	ls.lockedMessage = nil
	ls.Partition.ProcessBacklog(nil)
	return &pbpartition.AbortCommitResponse{}, nil
}

func (ls *ListenServer) Commit(ctx context.Context, req *pbpartition.CommitRequest) (*pbpartition.CommitResponse, error) {
	if ls.lockedMessage == nil {
		return nil, ErrNoLockedMessage
	}

	var err error
	if deleteMsg, ok := ls.lockedMessage.(*prototypes.DeleteRequest); ok {
		_, err = ls.Delete(ctx, deleteMsg)
	} else if setMsg, ok := ls.lockedMessage.(*prototypes.SetRequest); ok {
		_, err = ls.Set(ctx, setMsg)
	} else {
		return nil, ErrUnsupported2PCMsg
	}

	if err != nil {
		return nil, err
	}

	ls.lockedMessage = nil
	ls.Partition.ProcessBacklog(nil) // TODO: consider removing processing of backlog after 2PC termination
	return &pbpartition.CommitResponse{}, nil
}
