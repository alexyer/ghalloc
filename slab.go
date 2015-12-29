package ghalloc

type slabClass struct {
	ChunkSize int // Chunk size of the given class.
	SlabSize  int
	slabs     []*slab // Array of slabs of the given class.
}

type slab struct {
	memory []byte // Allocated memory of the given slab.
}

func newSlabClass(chunkSize, slabSize int) *slabClass {
	return &slabClass{
		ChunkSize: chunkSize,
		SlabSize:  slabSize,
		slabs:     []*slab{newSlab(slabSize)},
	}
}

func newSlab(size int) *slab {
	return &slab{
		memory: make([]byte, size),
	}
}

// Allocate new slab.
func (s *slabClass) grow() {
	s.slabs = append(s.slabs, newSlab(s.SlabSize))
}
