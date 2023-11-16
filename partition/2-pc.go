package partition

import (
	"context"

	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

func (ls *ListenServer) PrepareCommit(ctx context.Context, req *pbpartition.PrepareCommitRequest) (*pbpartition.PrepareCommitResponse, error) {
	if deleteMsg := req.GetDeleteRequest(); deleteMsg != nil {
		ls.lockedMessage = deleteMsg
	} else if setMsg := req.GetSetRequest(); setMsg != nil {
		ls.lockedMessage = setMsg
	} else {
		return nil, ErrUnsupported2PCMsg
	}

	// lock a db until locked message is committed or aborted
	ls.isLocked = true
	return &pbpartition.PrepareCommitResponse{}, nil
}
