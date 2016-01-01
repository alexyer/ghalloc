package ghalloc

import (
	"testing"
	"unsafe"
)

func TestNewSlab(t *testing.T) {
	sc := newSlabClass(512*KB, 1*MB)
	slab := newSlab(sc)

	if len(slab.memory) != 1*MB {
		t.Fatal("slab: allocated wrong memory slice")
	}
}

func TestAllocChunk(t *testing.T) {
	sc := newSlabClass(512*KB, 1*MB)
	slab := newSlab(sc)

	ptr := slab.allocChunk()

	if ptr == nil {
		t.Fatal("slab alloc: allocated nil")
	}

	if slab.allocated != 1 {
		t.Fatal("slab alloc: wrong allocation counter")
	}

	ptr = slab.allocChunk()

	if ptr == nil {
		t.Fatal("slab alloc: allocated nil")
	}

	if !slab.full {
		t.Fatal("slab alloc: wrong slab status. Expected full.")
	}

	ptr = slab.allocChunk()

	if ptr != nil {
		t.Fatal("slab alloc: Expected nil, got pointer")
	}
}

func TestFreeChunk(t *testing.T) {
	var (
		sc         *slabClass = newSlabClass(512*KB, 1*MB)
		slab       *slab      = newSlab(sc)
		ptr        unsafe.Pointer
		invalidPtr unsafe.Pointer = unsafe.Pointer(&sc)
	)

	for !slab.full {
		ptr = slab.allocChunk()
	}

	slab.freeChunk(invalidPtr)

	if !slab.full {
		t.Fatal("slab free chunk: freed wrong chunk")
	}

	slab.freeChunk(ptr)

	if slab.full {
		t.Fatal("slab free chunk: chunk is not freed")
	}
}
