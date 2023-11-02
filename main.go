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

		domain := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil) // domain of SHA-256

		percentInt, err := strconv.Atoi(args[4])
		if err != nil {
			panic(err)
		}
		var from *big.Int
		if percentInt == 0 {
			from = big.NewInt(0)
		} else {
			// from = domain * percentInt / 100
			from = new(big.Int).Div(
				new(big.Int).Mul(domain, big.NewInt(int64(percentInt))),
				big.NewInt(100),
			)
		}

		percentInt, err = strconv.Atoi(args[5])
		if err != nil {
			panic(err)
		}

		to := new(big.Int).Div(
			new(big.Int).Mul(domain, big.NewInt(int64(percentInt))),
			big.NewInt(100),
		)

		partition.RunPartitionServer(int64(port), dbPath, from, to)
	}
}
