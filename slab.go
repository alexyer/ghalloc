package ghalloc

type slabClass struct {
	ChunkSize int // Chunk size of the given class.
	SlabSize  int
	slabs     []*slab // Array of slabs of the given class.
}

type slab struct {
}

func NewSlabClass(chunkSize, slabSize int) *slabClass {
	return &slabClass{
		ChunkSize: chunkSize,
		SlabSize:  slabSize,
		slabs:     []*slab{&slab{}},
	}
}

// Allocate new slab.
func (s *slabClass) grow() {
	s.slabs = append(s.slabs, &slab{})
}
