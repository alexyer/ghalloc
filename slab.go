package ghalloc

import "unsafe"

type slab struct {
	slabClass *slabClass // Slab class of the slab.
	memory    []byte     // Allocated memory of the slab.
	full      bool       // Indicates if slab has free chunks to allocate.
	allocated uint64     // Number of allocated chunks. Also pointer to the first free chunk in the slab.
	chunkMap  uint64     // Bitmap of allocated chunks.
}

func newSlab(sc *slabClass) *slab {
	s := &slab{
		slabClass: sc,
		memory:    make([]byte, sc.SlabSize),
		full:      false,
		allocated: 0,
	}

	return s
}

// Allocate new pointer in the chunk.
func (s *slab) allocChunk() unsafe.Pointer {
	if s.full {
		return nil
	}

	i := s.getUnusedChunkIndex()
	ptr := unsafe.Pointer(&s.memory[s.getUnusedChunkIndex()])

	s.chunkMap |= 1 << i
	s.allocated++

	if s.allocated >= s.slabClass.Capacity {
		s.full = true
	}

	return ptr
}

// Free allocated pointer.
func (s *slab) freeChunk(ptr unsafe.Pointer) {
	uptr := uintptr(ptr)
	begin := uintptr(unsafe.Pointer(&s.memory[0]))

	if uptr >= begin && uptr <= uintptr(unsafe.Pointer(&s.memory[0]))+uintptr(s.slabClass.SlabSize) {
		s.chunkMap &^= 1 << uint64(uptr-begin) % uint64(s.slabClass.SlabSize)

		s.allocated--
	}
}

func (s *slab) getUnusedChunkIndex() uint64 {
	var i uint64

	for i = 0; i < s.slabClass.Capacity; i++ {
		if s.chunkMap&(1<<i) == 0 {
			return i
		}
	}
	// Just for complier.
	// Used if slab is not full, hence there are always free chunks.
	return i
}
