# Go Locks

[![CI](https://github.com/DaanV2/go-locks/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/DaanV2/go-locks/actions/workflows/pipeline.yaml)

A simple library that provides pools of locks for Go. It is useful when you need to lock on a resource that cannot carry its lock.
Such as files, network connections, etc.

## Usage
```go
package main

import (
    "fmt"
    "github.com/DaanV2/go-locks"
)

func main() {
    pool := locks.NewPool(100)

    lock := pool.GetLock(uint64)
    lock.Lock()
    defer lock.Unlock()

    // Do something with the resource

    // For files:
    key := locks.KeyForString("file.txt")
    lock := pool.GetLock(key)
    lock.Lock()
    defer lock.Unlock()
}
```