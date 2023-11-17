package balancer

import (
	"context"
	"fmt"
	"sync"

	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

func (b *Balancer) SetAtomic(ctx context.Context, key string, value []byte) error {
	range_, err := b.getRangeFromDigest([]byte(key))
	if err != nil {
		return err
	}

	clients := b.clients[*range_]
	if len(clients) == 0 {
		return ErrRangeNotYetCovered
	}

	err = b.prepareCommit(clients, key, value)
	// if >= 1 partition aborted, abort all
	if err != nil {
		b.abortCommit(ctx, clients)
	} else {
		b.commit(ctx, clients)
	}

	return nil
}

// prepareCommit sends a prepare commit request to all partitions that are responsible for the given key and awaits for their responses.
func (b *Balancer) prepareCommit(partitionClients []pbpartition.PartitionServiceClient, key string, value []byte) error {
	var wg sync.WaitGroup
	channel := make(chan error, len(partitionClients))
	for _, client := range partitionClients {
		wg.Add(1)
		clientCopy := client
		go func() {
			resp, err := clientCopy.PrepareCommit(context.Background(), &pbpartition.PrepareCommitRequest{})
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

func (b *Balancer) commit(ctx context.Context, partitionClients []pbpartition.PartitionServiceClient) error {
	var wg sync.WaitGroup
	channel := make(chan error, len(partitionClients))
	for _, client := range partitionClients {
		wg.Add(1)
		clientCopy := client
		go func() {
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
func (b *Balancer) abortCommit(ctx context.Context, partitionClients []pbpartition.PartitionServiceClient) {
	var wg sync.WaitGroup
	channel := make(chan error, len(partitionClients))
	for _, client := range partitionClients {
		wg.Add(1)
		clientCopy := client
		go func() {
			_, err := clientCopy.AbortCommit(ctx, &pbpartition.AbortCommitRequest{})
			if err != nil {
				// TODO: some error handling here
				fmt.Println(err, "TODO: Unimplemented branch")
			}

			channel <- nil
		}()
	}

	wg.Wait()

	for i := 0; i < len(partitionClients); i++ {
		if <-channel != nil {
			fmt.Println("TODO: Unimplemented branch")
			return
		}
	}
}
