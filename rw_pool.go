package locks

import "sync"

type RWPool struct {
	locks []*sync.RWMutex
}

// NewRWPool creates a new pool of locks, where amount is the number of locks to create.
// Instead of mutexes, it uses RWMutexes.
func NewRWPool(amount int) *RWPool {
	if amount <= 0 {
		amount = 10
	}

	locks := make([]*sync.RWMutex, 0, amount)
	for i := 0; i < amount; i++ {
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

// Len returns the amount of locks in the pool.
func (p *RWPool) Len() int {
	return len(p.locks)
}
