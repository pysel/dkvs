package partition_test

import (
	"crypto/sha256"
	"math/big"
	"os"
	"testing"

	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/testutil"
	hashrange "github.com/pysel/dkvs/types/hashrange"
	"github.com/stretchr/testify/require"
)

// Half of MaxInt
var defaultHashrange = hashrange.NewRange(big.NewInt(0).Bytes(), new(big.Int).Div(hashrange.MaxInt, big.NewInt(2)).Bytes())

func TestDatabaseMethods(t *testing.T) {
	p := partition.NewPartition("test")
	p.SetHashrange(defaultHashrange)

	defer p.Close()
	defer require.NoError(t, os.RemoveAll("test"))

	err := p.Set(testutil.DomainKey[:], []byte("Value"))
	require.Error(t, err, "should return error if key is not 32 bytes long - not a valid SHA-2 digest")

	hashedPartitionKey := sha256.Sum256(testutil.DomainKey)
	hashedNotPartitionKey := sha256.Sum256(testutil.NonDomainKey)

	err = p.Set(hashedPartitionKey[:], []byte("Value"))
	require.NoError(t, err) // partition's key, should store correctly

	err = p.Set(hashedNotPartitionKey[:], []byte("Value2"))
	require.Error(t, err) // not partition's key, should return error

	value, err := p.Get(hashedPartitionKey[:])
	require.NoError(t, err) // partition's key, should get correctly
	require.Equal(t, []byte("Value"), value)

	value, err = p.Get(hashedNotPartitionKey[:])
	require.Error(t, err) // not partition's key, should return error
	require.Nil(t, value)

	err = p.Delete(hashedNotPartitionKey[:])
	require.Error(t, err) // not partition's key, should return error

	err = p.Delete(hashedPartitionKey[:])
	require.NoError(t, err) // partition's key, should delete correctly

	value, err = p.Get(hashedPartitionKey[:])
	require.NoError(t, err) // partition's key, should return error
	require.Nil(t, value)
}
