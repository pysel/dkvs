package partition

import (
	"context"
	"math/big"

	"github.com/pysel/dkvs/prototypes"
	"github.com/pysel/dkvs/shared"
	"github.com/pysel/dkvs/types"
	"google.golang.org/protobuf/proto"
)

// Set sets a value for a key.
func (ls *ListenServer) Set(ctx context.Context, req *prototypes.SetRequest) (resp *prototypes.SetResponse, err error) {
	defer func() { ls.postCRUD(err, req.String()) }()

	// note: if request is not valid, the timestamp will not be incremented
	// TODO: investigate if it is a valid behaviour.
	if err = req.Validate(); err != nil {
		return nil, err
	}

	// process logical timestamp
	if err = ls.validateTSGrpcLevel(req.Lamport, req); err != nil {
		return nil, err
	}

	var value []byte
	value, err = reqToBytes(req)
	if err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)
	err = ls.Partition.Set(shaKey[:], value)
	if err != nil {
		return nil, ErrInternal{Reason: err}
	}

	// Log event
	ls.EventHandler.Emit(SetEvent{key: string(req.Key), data: string(req.Value)})

	return &prototypes.SetResponse{}, nil
}

// Get gets a value for a key.
func (ls *ListenServer) Get(ctx context.Context, req *prototypes.GetRequest) (resp *prototypes.GetResponse, err error) {
	defer func() { ls.postCRUD(err, req.String()) }()

	if err = req.Validate(); err != nil {
		return nil, err
	}

	// process logical timestamp
	if err = ls.validateTSGrpcLevel(req.Lamport, req); err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)

	value, err := ls.Partition.Get(shaKey[:])
	if err != nil {
		return nil, ErrInternal{Reason: err}
	}
	if value == nil {
		return &prototypes.GetResponse{StoredValue: nil}, nil
	}

	var storedValue prototypes.StoredValue
	err = proto.Unmarshal(value, &storedValue)
	if err != nil {
		return nil, err
	}

	// Log event
	ls.EventHandler.Emit(GetEvent{key: string(req.Key), returned: string(storedValue.Value)})

	return &prototypes.GetResponse{StoredValue: &storedValue}, nil
}

// Delete deletes a value for a key.
func (ls *ListenServer) Delete(ctx context.Context, req *prototypes.DeleteRequest) (resp *prototypes.DeleteResponse, err error) {
	defer func() { ls.postCRUD(err, req.String()) }()

	if err = req.Validate(); err != nil {
		return nil, err
	}

	// process logical timestamp
	if err = ls.validateTSGrpcLevel(req.Lamport, req); err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)

	err = ls.Partition.Delete(shaKey[:])
	if err != nil {
		return nil, ErrInternal{Reason: err}
	}

	// Log event
	ls.EventHandler.Emit(DeleteEvent{key: string(req.Key)})

	return &prototypes.DeleteResponse{}, nil
}

// SetHashrange sets the hashrange for this partition.
func (ls *ListenServer) SetHashrange(ctx context.Context, req *prototypes.SetHashrangeRequest) (res *prototypes.SetHashrangeResponse, err error) {
	defer func() {
		if err != nil {
			ls.EventHandler.Emit(shared.ErrorEvent{Req: req.String(), Err: err})
		} else {
			min := new(big.Int).SetBytes(req.Min)
			max := new(big.Int).SetBytes(req.Max)

			ls.EventHandler.Emit(SetHashrangeEvent{min: min, max: max})
		}
	}()

	if req == nil {
		err = types.ErrNilRequest
		return
	}

	ls.hashrange = NewRange(req.Min, req.Max)
	return &prototypes.SetHashrangeResponse{}, nil
}

// postCRUD runs functionality that should be run after every CRUD operation.
func (ls *ListenServer) postCRUD(err error, req string) {
	// if error is a warning, log it as warning
	// log error as error otherwise
	if err != nil {
		if eventError, ok := err.(shared.IsWarningEventError); ok {
			ls.EventHandler.Emit(eventError.WarningErrorToEvent(req))
			return
		}
		ls.EventHandler.Emit(shared.ErrorEvent{Req: req, Err: err})
	} else {
		ls.IncrTs()
	}

	ls.ProcessBacklog()
}

func (ls *ListenServer) validateTSGrpcLevel(ts uint64, message proto.Message) error {
	switch err := ls.Partition.validateTS(ts); err.(type) {
	case ErrTimestampIsStale: // stale/already processed request
		return err
	case ErrTimestampNotNext: // replica is not ready to process this request
		ls.backlog.Add(ts, message)
		return err // let balancer know that this replica is not ready for the request
	}

	return nil
}
