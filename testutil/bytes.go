package testutil

func ReverseByteOrder(data []byte) []byte {
	// Get the length of the data
	length := len(data)

	// Swap the elements to reverse the byte order
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		data[i], data[j] = data[j], data[i]
	}

	return data
}
