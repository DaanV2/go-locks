package locks

import (
	"crypto/sha1"
	"encoding/binary"
)

// KeyForString returns a key for a string
func KeyForString(s string) uint64 {
	return KeyForBytes([]byte(s))
}

// KeyForBytes returns a key for a byte slice
func KeyForBytes(b []byte) uint64 {
	data := sha1.Sum(b)

	return binary.BigEndian.Uint64(data[:])
}