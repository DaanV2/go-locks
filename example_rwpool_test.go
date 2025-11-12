package locks_test

import (
	"fmt"
	"sync"

	"github.com/daanv2/go-locks"
)

// ExampleRWPool demonstrates basic usage of a read-write lock pool
func ExampleRWPool() {
	// Create a read-write lock pool with default size
	pool := locks.NewRWPool()

	// Get a lock for a specific resource
	resourceID := uint64(42)
	lock := pool.GetLock(resourceID)

	// Write lock
	lock.Lock()
	fmt.Println("Write lock acquired")
	lock.Unlock()

	// Read lock
	lock.RLock()
	fmt.Println("Read lock acquired")
	lock.RUnlock()

	// Output:
	// Write lock acquired
	// Read lock acquired
}

// ExampleRWPool_multipleReaders demonstrates multiple concurrent readers
func ExampleRWPool_multipleReaders() {
	pool := locks.NewRWPool()
	var wg sync.WaitGroup

	cacheKey := "user:1234"
	
	// Multiple readers can access simultaneously
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			lock := pool.GetLockByString(cacheKey)
			lock.RLock()
			defer lock.RUnlock()
			// Read operation - multiple readers can proceed concurrently
		}(i)
	}

	wg.Wait()
	fmt.Println("All readers completed")

	// Output:
	// All readers completed
}

// ExampleRWPool_GetLockByString demonstrates read-write locking with string keys
func ExampleRWPool_GetLockByString() {
	pool := locks.NewRWPool()

	configFile := "config.json"
	lock := pool.GetLockByString(configFile)

	// Multiple readers
	lock.RLock()
	fmt.Println("Reading configuration")
	lock.RUnlock()

	// Single writer
	lock.Lock()
	fmt.Println("Updating configuration")
	lock.Unlock()

	// Output:
	// Reading configuration
	// Updating configuration
}

// ExampleNewRWPool demonstrates creating an RWPool with custom size
func ExampleNewRWPool() {
	// Create with default size
	defaultPool := locks.NewRWPool()
	fmt.Printf("Default RWPool size: %d\n", defaultPool.Len())

	// Create with custom size
	customPool := locks.NewRWPool(locks.WithSize(100))
	fmt.Printf("Custom RWPool size: %d\n", customPool.Len())

	// Output:
	// Default RWPool size: 40
	// Custom RWPool size: 100
}
