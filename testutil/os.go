package testutil

import "os"

func RemovePaths(paths []string) {
	for _, path := range paths {
		os.RemoveAll(path)
	}
}
