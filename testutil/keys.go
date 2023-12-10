package testutil

var (
	// DomainKey is a key that is in the domain of first half of sha-2 domain.
	DomainKey = []byte("Partition key")
	// NonDomainKey is a key that is in the domain of second half of sha-2 domain.
	NonDomainKey = []byte("Not partition key.")
)
