package partition_test

import (
	"crypto/sha256"
	"os"
	"testing"

	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/testutil"
	"github.com/stretchr/testify/require"
)

func TestPartitionEvents(t *testing.T) {
	defer os.RemoveAll(testutil.TestDBPath)

	p := partition.NewPartition(testutil.TestDBPath)
	defer p.Close()

	p.SetHashrange(testutil.DefaultHashrange)

	key := sha256.Sum256(testutil.DomainKey)

	// Set
	err := p.Set(key[:], []byte("Value"))
	require.NoError(t, err)
}
