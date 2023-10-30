package partition

import (
	"crypto/sha256"
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Half of maxInt
var defaultHashRange = NewRange(big.NewInt(0), maxInt.Div(maxInt, big.NewInt(2)))

func TestDatabaseMethods(t *testing.T) {
	p := NewPartition("test", defaultHashRange)
	defer p.Close()
	defer require.NoError(t, os.RemoveAll("test"))

	partitionKey := []byte("Partition key")
	notPartitionKey := []byte("Not partition key.")

	hashedPartitionKey := sha256.Sum256(partitionKey)
	hashedNotPartitionKey := sha256.Sum256(notPartitionKey)

	err := p.Set(hashedPartitionKey[:], []byte("Value"))
	require.NoError(t, err) // partition's key, should store correctly

	err = p.Set(hashedNotPartitionKey[:], []byte("Value2"))
	require.Error(t, err) // not partition's key, should return error

	value, err := p.Get(hashedPartitionKey[:])
	require.NoError(t, err) // partition's key, should get correctly
	require.Equal(t, []byte("Value"), value)

	value, err = p.Get(hashedNotPartitionKey[:])
	require.Error(t, err) // not partition's key, should return error
	require.Nil(t, value)

	has, err := p.Has(hashedPartitionKey[:])
	require.NoError(t, err)
	require.True(t, has) // partition's key, should return true

	has, err = p.Has(hashedNotPartitionKey[:])
	require.Error(t, err)
	require.False(t, has) // not partition's key, should return false

	err = p.Delete(hashedNotPartitionKey[:])
	require.Error(t, err) // not partition's key, should return error

	err = p.Delete(hashedPartitionKey[:])
	require.NoError(t, err) // partition's key, should delete correctly

	value, err = p.Get(hashedPartitionKey[:])
	require.Error(t, err) // partition's key, should return error
	require.Equal(t, value, []byte{})
}
