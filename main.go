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
		fmt.Println(args)
		if len(args) < 5 {
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

		percentInt, err := strconv.Atoi(args[5])
		if err != nil {
			panic(err)
		}

		domain := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)

		to := new(big.Int).Div(domain, big.NewInt(int64(100/percentInt)))
		fmt.Println(to.Bytes(), len(to.Bytes()))

		// to, ok := new(big.Int).SetString(args[5], 10)
		// if !ok {
		// 	panic("Could not parse to")
		// }

		partition.RunPartitionServer(int64(port), dbPath, from, to)
	}
}
