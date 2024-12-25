package locks

import "sync"

// RWPool is a Read Write Pool of locks
type RWPool struct {
	locks []*sync.RWMutex
}

// NewRWPool creates a new pool of locks, where amount is the number of locks to create.
// Instead of mutexes, it uses RWMutexes.
func NewRWPool(opts ...PoolOption) *RWPool {
	opt := DefaultOptions()
	opt.Modify(opts...)
	opt.Sanitize()

	locks := make([]*sync.RWMutex, 0, opt.Size)
	for range opt.Size {
		locks = append(locks, new(sync.RWMutex))
	}
	return &RWPool{locks}
}

// GetLock returns a lock from the pool based on the key.
// The key is provided to ensure that the same lock is always returned for the same key. And can be any value.
// Example:
// 	lock := pool.GetLock(1)
// 	lock = pool.GetLock(987654321)
func (p *RWPool) GetLock(key uint64) *sync.RWMutex {
	length := uint64(len(p.locks))
	index := key % length

	return p.locks[index]
}

// GetLockByString returns a key for a string
func (p *RWPool) GetLockByString(key string) *sync.RWMutex {
	id := KeyForString(key)

	return p.GetLock(id)
}

// GetLockByBytes returns a key for a byte slice
func (p *RWPool) GetLockByBytes(b []byte) *sync.RWMutex {
	id := KeyForBytes(b)

	return p.GetLock(id)
}

// Len returns the amount of locks in the pool.
func (p *RWPool) Len() int {
	return len(p.locks)
}
