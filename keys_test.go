package locks_test

import (
	"testing"

	"github.com/daanv2/go-locks"
	"github.com/stretchr/testify/require"
)

func Test_KeyForString(t *testing.T) {
	k := locks.KeyForString("hello world")
	
	require.NotEqual(t, uint64(0), k)
}

func Fuzz_KeyForString(f *testing.F) {
	f.Add("hello world")
	f.Add("PY9S/4liWz3kI8t1Wjm6QQ==")
	f.Add("dir/dir/file.txt")

	f.Fuzz(func(t *testing.T, s string) {
		k := locks.KeyForString(s)
	
		require.NotEqual(t, uint64(0), k)
	})
}