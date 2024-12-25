package locks

import (
	"sync"
)

// Pool is a pool of locks that can be used to lock based on a key.
type Pool struct {
	locks []*sync.Mutex
}

// NewPool creates a new pool of locks, where amount is the number of locks to create.
// It returns an error if the amount is less than or equal to 0.
// The ideal amount of locks is dependent on the number of threads. With a factor that depends on the amount of collision you want per lock.
// Example:
//
//	threads := 7
//	// Collision of 25% per lock
//	amount := threads * (100 / 25) // => 7 * 4 = 28
func NewPool(opts ...PoolOption) *Pool {
	opt := DefaultOptions()
	opt.Modify(opts...)
	opt.Sanitize()

	locks := make([]*sync.Mutex, 0, opt.Size)
	for range opt.Size {
		locks = append(locks, new(sync.Mutex))
	}
	return &Pool{locks}
}

// GetLock returns a lock from the pool based on the key.
// The key is provided to ensure that the same lock is always returned for the same key. And can be any value.
// Example:
//
//	lock := pool.GetLock(1)
//	lock = pool.GetLock(987654321)
func (p *Pool) GetLock(key uint64) *sync.Mutex {
	length := uint64(len(p.locks))
	index := key % length

	return p.locks[index]
}

// GetLockByString returns a key for a string
func (p *Pool) GetLockByString(key string) *sync.Mutex {
	id := KeyForString(key)

	return p.GetLock(id)
}

// GetLockByBytes returns a key for a byte slice
func (p *Pool) GetLockByBytes(b []byte) *sync.Mutex {
	id := KeyForBytes(b)

	return p.GetLock(id)
}

// Len returns the amount of locks in the pool.
func (p *Pool) Len() int {
	return len(p.locks)
}
