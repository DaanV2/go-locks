package locks_test

import (
	"fmt"

	"github.com/daanv2/go-locks"
)

// ExampleKeyForString demonstrates generating a key from a string
func ExampleKeyForString() {
	// Generate a key for a string identifier
	key := locks.KeyForString("user:12345")
	
	fmt.Printf("Key generated: %t\n", key != 0)

	// Output:
	// Key generated: true
}

// ExampleKeyForBytes demonstrates generating a key from bytes
func ExampleKeyForBytes() {
	data := []byte("important-resource")
	key := locks.KeyForBytes(data)
	
	fmt.Printf("Key generated: %t\n", key != 0)

	// Output:
	// Key generated: true
}

// ExampleKeyForString_consistency demonstrates that the same string always produces the same key
func ExampleKeyForString_consistency() {
	resource := "shared-resource"
	
	key1 := locks.KeyForString(resource)
	key2 := locks.KeyForString(resource)
	
	fmt.Printf("Keys are consistent: %t\n", key1 == key2)

	// Output:
	// Keys are consistent: true
}
