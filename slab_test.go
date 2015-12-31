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

	if len(slab.chunks) != 2 {
		t.Fatal("slab: wrong slab memory partition")
	}

	if slab.chunks[0] != unsafe.Pointer(&slab.memory[0]) && slab.chunks[1] != unsafe.Pointer(uintptr(unsafe.Pointer(&slab.memory[0]))+uintptr(512*KB)) {
		t.Fatal("slab: wrong pointers")
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
func TestFindChunk(t *testing.T) {
	sc := newSlabClass(512*KB, 1*MB)
	slab := newSlab(sc)

	ptr := slab.allocChunk()
	invalidPtr := unsafe.Pointer(&sc)

	if slab.findChunk(ptr) < 0 {
		t.Fatal("slab find chunk: chunk is not found")
	}

	if slab.findChunk(invalidPtr) >= 0 {
		t.Fatal("slab find chunk: wrong chunk found")
	}
}

func TestFreeChunk(t *testing.T) {
	var (
		sc   *slabClass = newSlabClass(512*KB, 1*MB)
		slab *slab      = newSlab(sc)
		ptr  unsafe.Pointer
	)

	for !slab.full {
		ptr = slab.allocChunk()
	}

	slab.freeChunk(slab.chunks[0])

	if slab.full {
		t.Fatal("slab free chunk: chunk is not freed")
	}

	slab.freeChunk(ptr)

	if slab.full {
		t.Fatal("slab free chunk: chunk is not freed")
	}
}

func TestAllocChunkRaces(t *testing.T) {
	sc := newSlabClass(512*KB, 1*MB)
	slab := newSlab(sc)

	go slab.allocChunk()
	go slab.allocChunk()
}

func TestFreeChunkRaces(t *testing.T) {
	sc := newSlabClass(512*KB, 1*MB)
	slab := newSlab(sc)

	ptr1 := slab.allocChunk()
	ptr2 := slab.allocChunk()

	go slab.freeChunk(ptr1)
	go slab.freeChunk(ptr2)
}
