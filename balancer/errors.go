package balancer

import "errors"

var (
	ErrPartitionOverflow        = errors.New("enough partitions are already registered")
	ErrCoverageNotProperlySetUp = errors.New("coverage is not properly set up")
	ErrDigestNotCovered         = errors.New("digest is not covered by any range")
	ErrRangeNotYetCovered       = errors.New("range is not yet covered by any partition")
	ErrAllReplicasFailed        = errors.New("all replicas failed to process request")

	// 2PC
	ErrPrepareCommitAborted = errors.New("prepare commit aborted")
	ErrCommitAborted        = errors.New("commit aborted")
)

type ErrDecisionNotSavedToDisk struct {
	Reason   error
	Decision []byte
}

func (e ErrDecisionNotSavedToDisk) Error() string {
	return "decision not saved to disk: " + e.Reason.Error()
}

func (e ErrDecisionNotSavedToDisk) Unwrap() error {
	return e.Reason
}

type ErrDecisionWasNotCleared struct {
	Reason error
}

func (e ErrDecisionWasNotCleared) Error() string {
	return "decision was not cleared from disk after a two-phase commit has ended: " + e.Reason.Error()
}

func (e ErrDecisionWasNotCleared) Unwrap() error {
	return e.Reason
}
