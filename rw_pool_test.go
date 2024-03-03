package locks_test

import (
	"math"
	"testing"

	"github.com/DaanV2/go-locks"
	"github.com/stretchr/testify/require"
)

func Test_RWPool_GetLock(t *testing.T) {
	t.Run("Simple get works", func(t *testing.T) {
		pool, err := locks.NewRWPool(10)
		require.NoError(t, err)

		lock := pool.GetLock(0)
		require.NotNil(t, lock)
	})

	t.Run("Key greater than size works", func(t *testing.T) {
		pool, err := locks.NewRWPool(10)
		require.NoError(t, err)

		lock := pool.GetLock(100)
		require.NotNil(t, lock)
	})

	t.Run("Large size works", func(t *testing.T) {
		pool, err := locks.NewRWPool(10)
		require.NoError(t, err)

		lock := pool.GetLock(987654321)
		require.NotNil(t, lock)
	})

	t.Run("0 Amount should return an error", func(t *testing.T) {
		_, err := locks.NewRWPool(0)
		require.Error(t, err)
		require.Equal(t, err, locks.ErrInvalidAmount)
	})

	t.Run("Same key should return the same lock", func(t *testing.T) {
		pool, err := locks.NewRWPool(10)
		require.NoError(t, err)

		lock1 := pool.GetLock(987654321)
		lock2 := pool.GetLock(987654321)
		require.Equal(t, lock1, lock2)
	})
}

func Fuzz_RWPool_GetLock(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(1))
	f.Add(uint64(100))
	f.Add(uint64(1000))
	f.Add(uint64(math.MaxUint64))

	f.Fuzz(func(t *testing.T, amount uint64) {
		pool, err := locks.NewRWPool(10)
		require.NoError(t, err)

		lock := pool.GetLock(amount)
		require.NotNil(t, lock)
	})
}
