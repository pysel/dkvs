package main

import (
	"fmt"
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
*/
func main() {
	args := os.Args

	if len(args) < 2 {
		panic("Provide a mode of this node")
	}
	mode := args[1]
	switch mode {
	case types.PARTITION_MODE:
		if len(args) < 3 {
			panic(fmt.Sprintf("Need a port, db path, but got: %T", args))
		}
		port, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}

		dbPath := args[3]

		partition.RunPartitionServer(int64(port), dbPath)
	}
}
