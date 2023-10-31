package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/types"
)

/*
If node is a partition, arguments should be (in mentioned order):
  - mode (partition)
  - port
  - db path
  - from (beginning of scope of keys for this partition)
  - to (end of scope of keys for this partition)
*/
func main() {
	args := os.Args

	if len(args) < 2 {
		panic("Provide a mode of this node")
	}
	mode := args[1]
	switch mode {
	case types.PARTITION_MODE:
		if len(args) < 6 {
			panic(fmt.Sprintf("Need a port, db path, from and to, but got: %T", args))
		}
		port, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}

		dbPath := args[3]

		from, ok := new(big.Int).SetString(args[4], 10)
		if !ok {
			panic("Could not parse from")
		}

		to, ok := new(big.Int).SetString(args[5], 10)
		if !ok {
			panic("Could not parse to")
		}

		partition.RunPartitionServer(int64(port), dbPath, from, to)
	}
}
