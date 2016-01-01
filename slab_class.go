package ghalloc

import "unsafe"

type slabClass struct {
	Capacity  uint64 // Number of chunks of the given class.
	ChunkSize int    // Chunk size of the given class.
	SlabSize  int
	slabs     []*slab // Array of slabs of the given class.
}

func newSlabClass(chunkSize, slabSize int) *slabClass {
	sc := &slabClass{
		Capacity:  uint64(slabSize / chunkSize),
		ChunkSize: chunkSize,
		SlabSize:  slabSize,
	}

	sc.slabs = []*slab{newSlab(sc)}

	return sc
}

// Allocate new slab.
func (s *slabClass) grow() *slab {
	ns := newSlab(s)
	s.slabs = append(s.slabs, ns)
	return ns
}

// Get pointer to a free chunk in the given slab class.
func (s *slabClass) getChunk() unsafe.Pointer {
	slab := s.findAvailableSlab()

	if slab == nil {
		slab = s.grow()
	}

	return slab.allocChunk()
}

// Return pointer to a free chunk list.
func (s *slabClass) returnChunk(ptr unsafe.Pointer) {
	uptr := uintptr(ptr)
	for i := 0; i < len(s.slabs); i++ {
		if uptr <= uintptr(unsafe.Pointer(&s.slabs[i].memory[0]))+uintptr(s.SlabSize) {
			s.slabs[i].freeChunk(ptr)
		}
	}
}

// Find non full slab.
func (s *slabClass) findAvailableSlab() *slab {
	for i := 0; i < len(s.slabs); i++ {
		if !s.slabs[i].full {
			return s.slabs[i]
		}
	}
	return nil
}
