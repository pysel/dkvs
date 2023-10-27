package partition

import "bytes"

func (p *Partition) checkKeyRange(key []byte) error {
	if bytes.Compare(key, p.hashrange.Min.Bytes()) == -1 || bytes.Compare(key, p.hashrange.Max.Bytes()) == 1 {
		// Key is out of range.
		return ErrNotThisPartitionKey
	}

	return nil
}
