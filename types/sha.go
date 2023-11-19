package types

import "crypto/sha256"

func ShaKey(key string) [32]byte {
	checksum := sha256.Sum256([]byte(key))
	return checksum
}
