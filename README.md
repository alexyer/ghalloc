[![Build Status](https://travis-ci.org/alexyer/ghalloc.svg)](https://travis-ci.org/alexyer/ghalloc)
[![Coverage Status](https://coveralls.io/repos/alexyer/ghalloc/badge.svg?branch=master&service=github)](https://coveralls.io/github/alexyer/ghalloc?branch=master)
[![GoDoc](https://godoc.org/github.com/alexyer/ghalloc?status.svg)](https://godoc.org/github.com/alexyer/ghalloc)

#Description
Go Slab Allocator.
The idea is similar to Memcached slab allocator.
Heavily uses unsafe package and pointer arithmetic, so use on your own risk.

## Example
```go
type SomeStructure struct {
    Field1 int
}

allocator := ghalloc.New(&ghalloc.Options{})
newStruct := (*SomeStructure)(ghalloc.Alloc(unsafe.Sizeof(SomeStructure{})))
newStruct.Field1 = 1
ghalloc.Free(&newStruct, unsafe.Sizeof(newStruct))
```
