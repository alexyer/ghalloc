package ghalloc

import (
	"sync"
	"unsafe"
)

type slab struct {
	slabMu    sync.Mutex
	slabClass *slabClass       // Slab class of the slab.
	memory    []byte           // Allocated memory of the slab.
	full      bool             // Indicates if slab has free chunks to allocate.
	allocated uint64           // Number of allocated chunks. Also pointer to the first free chunk in the slab.
	chunks    []unsafe.Pointer // Array of pointers to the chunks of the slab.
}

func newSlab(sc *slabClass) *slab {
	s := &slab{
		slabClass: sc,
		memory:    make([]byte, sc.SlabSize),
		full:      false,
		allocated: 0,
		chunks:    make([]unsafe.Pointer, sc.Capacity, sc.Capacity),
	}

	// Create array of chunks. Each chunk is a pointer to the memory region of the slab.
	var i uint64
	for i = 0; i < sc.Capacity; i++ {
		s.chunks[i] = unsafe.Pointer(uintptr(unsafe.Pointer(&s.memory[0])) + unsafe.Sizeof(s.memory[0])*uintptr(i))
	}

	return s
}

// Allocate new pointer in the chunk.
func (s *slab) allocChunk() unsafe.Pointer {
	if s.full {
		return nil
	}

	s.slabMu.Lock()

	ptr := s.chunks[s.allocated]
	s.allocated++

	if s.allocated >= s.slabClass.Capacity {
		s.full = true
	}

	s.slabMu.Unlock()

	return ptr
}

// Free allocated pointer.
func (s *slab) freeChunk(ptr unsafe.Pointer) {
	s.slabMu.Lock()
	i := s.findChunk(ptr)

	switch {
	// Free the last allocated chunk.
	// Just move border to the left.
	case uint64(i) == s.allocated-1:
		s.allocated--

		if s.full {
			s.full = false
		}

		s.slabMu.Unlock()
		return

	// Free chunk from the allocated range.
	// Copy the last allocated chunk to free chunk place
	// and move border to the left.
	case i >= 0:
		s.allocated--
		s.chunks[i] = s.chunks[s.allocated]

		if s.full {
			s.full = false
		}

		s.slabMu.Unlock()
		return

	// Does not belong to the current chunk
	default:
		s.slabMu.Unlock()
		return
	}
}

// Find ptr index in the array of chunk.
// Return index or -1 if chunk is not found.
func (s *slab) findChunk(ptr unsafe.Pointer) int {
	for i := 0; i < len(s.chunks); i++ {
		if ptr == s.chunks[i] {
			return i
		}
	}
	return -1
}
