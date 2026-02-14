package locks_test

import (
	"fmt"
	"sync"

	"github.com/daanv2/go-locks"
)

// ExamplePool demonstrates basic usage of a lock pool for resource locking
func ExamplePool() {
	// Create a pool with default size
	pool := locks.NewPool()

	// Get a lock for a specific resource ID
	resourceID := uint64(12345)
	lock := pool.GetLock(resourceID)

	// Use the lock
	lock.Lock()
	fmt.Println("Resource locked")
	lock.Unlock()
	fmt.Println("Resource unlocked")

	// Output:
	// Resource locked
	// Resource unlocked
}

// ExamplePool_GetLockByString demonstrates locking based on string keys, useful for file paths
func ExamplePool_GetLockByString() {
	pool := locks.NewPool()

	// Lock based on a file path
	filePath := "data/users.json"
	lock := pool.GetLockByString(filePath)

	lock.Lock()
	fmt.Println("File locked for writing")
	// ... perform file operations ...
	lock.Unlock()
	fmt.Println("File unlocked")

	// Output:
	// File locked for writing
	// File unlocked
}

// ExamplePool_concurrent demonstrates concurrent access using the same pool
func ExamplePool_concurrent() {
	pool := locks.NewPool()
	var wg sync.WaitGroup

	// Simulate multiple goroutines accessing different resources
	for i := range 3 {
		wg.Add(1)
		go func(id uint64) {
			defer wg.Done()
			lock := pool.GetLock(id)
			lock.Lock()
			defer lock.Unlock()
			// Critical section for this resource
		}(uint64(i))
	}

	wg.Wait()
	fmt.Println("All operations completed")

	// Output:
	// All operations completed
}

// ExampleNewPool demonstrates creating a pool with custom size
func ExampleNewPool() {
	// Create a pool with default size (size depends on GOMAXPROCS)
	defaultPool := locks.NewPool()
	fmt.Printf("Default pool has locks: %t\n", defaultPool.Len() > 0)

	// Create a pool with custom size
	customPool := locks.NewPool(locks.WithSize(50))
	fmt.Printf("Custom pool size: %d\n", customPool.Len())

	// Output:
	// Default pool has locks: true
	// Custom pool size: 50
}

// ExamplePool_GetLockByBytes demonstrates locking based on byte slices
func ExamplePool_GetLockByBytes() {
	pool := locks.NewPool()

	data := []byte("important-data")
	lock := pool.GetLockByBytes(data)

	lock.Lock()
	fmt.Println("Data locked")
	lock.Unlock()

	// Output:
	// Data locked
}
