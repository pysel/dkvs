package balancer

import (
	"context"
	"fmt"
	"sync"

	"github.com/pysel/dkvs/partition"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

// AtomicMessage sends a message to all partitions that are responsible for the given key and awaits for their responses.
// On successfull ack from all nodes, sends a commit message, else sends an abort message.
func (b *Balancer) AtomicMessage(ctx context.Context, range_ *partition.Range, msg *pbpartition.PrepareCommitRequest) error {
	clients := b.clients[range_]
	if len(clients) == 0 {
		return ErrRangeNotYetCovered
	}

	// synchronous prepare commit step
	err := b.prepareCommit(clients, msg)

	// If >= 1 partition aborted, abort all
	// Before aborting/committing, save decision to disk so that we can recover from a crash
	if err != nil {
		err := b.DB.Set(PrepareCommitDecisionKey, []byte("abort"))
		if err != nil {
			return ErrDecisionNotSavedToDisk{Reason: err, Decision: []byte("abort")}
		}

		b.abortCommit(ctx, clients)
	} else {
		err := b.DB.Set(PrepareCommitDecisionKey, []byte("commit"))
		if err != nil {
			return ErrDecisionNotSavedToDisk{Reason: err, Decision: []byte("commit")}
		}

		err = b.commit(ctx, clients)
		if err != nil {
			return err
		}
	}

	err = b.DB.Delete(PrepareCommitDecisionKey)
	if err != nil {
		return ErrDecisionWasNotCleared{Reason: err}
	}

	return nil
}

// prepareCommit sends a prepare commit request to all partitions that are responsible for the given key and awaits for their responses.
func (b *Balancer) prepareCommit(partitionClients []pbpartition.PartitionServiceClient, msg *pbpartition.PrepareCommitRequest) error {
	var wg sync.WaitGroup
	channel := make(chan error, len(partitionClients))
	for _, client := range partitionClients {
		wg.Add(1)
		clientCopy := client
		go func() {
			defer wg.Done()
			resp, err := clientCopy.PrepareCommit(context.Background(), msg)
			if err != nil {
				channel <- err
			}

			if resp != nil && resp.Ok {
				channel <- nil
			}

			if resp != nil && !resp.Ok {
				channel <- ErrPrepareCommitAborted
			}
		}()
	}

	wg.Wait()

	for i := 0; i < len(partitionClients); i++ {
		err := <-channel
		if err != nil {
			return err
		}
	}

	return nil
}

// commit sends a commit request to provided partitions.
func (b *Balancer) commit(ctx context.Context, partitionClients []pbpartition.PartitionServiceClient) error {
	var wg sync.WaitGroup
	channel := make(chan error, len(partitionClients))
	for _, client := range partitionClients {
		wg.Add(1)
		clientCopy := client
		go func() {
			defer wg.Done()
			_, err := clientCopy.Commit(ctx, &pbpartition.CommitRequest{})
			if err != nil {
				channel <- err
			} else {
				channel <- nil
			}
		}()
	}

	wg.Wait()

	for i := 0; i < len(partitionClients); i++ {
		if <-channel != nil {
			return ErrCommitAborted
		}
	}

	return nil
}

// abortCommit sends an abort commit request to provided partitions.
func (b *Balancer) abortCommit(ctx context.Context, partitionClients []pbpartition.PartitionServiceClient) {
	var wg sync.WaitGroup
	channel := make(chan error, len(partitionClients))
	for _, client := range partitionClients {
		wg.Add(1)
		clientCopy := client
		go func() {
			defer wg.Done()
			_, err := clientCopy.AbortCommit(ctx, &pbpartition.AbortCommitRequest{})
			if err != nil {
				// TODO: some error handling here
				fmt.Println(err, "TODO: Unimplemented branch 1")
			}

			channel <- nil
		}()
	}

	wg.Wait()

	for i := 0; i < len(partitionClients); i++ {
		if <-channel != nil {
			fmt.Println("TODO: Unimplemented branch 2")
			return
		}
	}
}
