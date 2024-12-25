package locks_test

import (
	"math"
	"testing"

	"github.com/daanv2/go-locks"
	"github.com/stretchr/testify/require"
)

func Test_Pool_GetLock(t *testing.T) {
	t.Run("Simple get works", func(t *testing.T) {
		pool := locks.NewPool(10)

		lock := pool.GetLock(0)
		require.NotNil(t, lock)
	})

	t.Run("Key greater than size works", func(t *testing.T) {
		pool := locks.NewPool(10)

		lock := pool.GetLock(100)
		require.NotNil(t, lock)
	})

	t.Run("Large size works", func(t *testing.T) {
		pool := locks.NewPool(10)

		lock := pool.GetLock(987654321)
		require.NotNil(t, lock)
	})

	t.Run("0 Amount should return a pool with items", func(t *testing.T) {
		pool := locks.NewPool(0)
		require.NotNil(t, pool)
		l := pool.Len()
		require.Greater(t, l, 0)
	})

	t.Run("Same key should return the same lock", func(t *testing.T) {
		pool := locks.NewPool(10)

		lock1 := pool.GetLock(987654321)
		lock2 := pool.GetLock(987654321)
		require.Equal(t, lock1, lock2)
	})
}

func Fuzz_Pool_GetLock(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(1))
	f.Add(uint64(100))
	f.Add(uint64(1000))
	f.Add(uint64(math.MaxUint64))

	f.Fuzz(func(t *testing.T, amount uint64) {
		pool := locks.NewPool(10)

		lock := pool.GetLock(amount)
		require.NotNil(t, lock)
	})
}
