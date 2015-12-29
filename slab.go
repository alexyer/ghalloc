package ghalloc

type slabClass struct {
	ChunkSize int     // Chunk size of the given class.
	slabs     []*slab // Array of slabs of the given class.
}

type slab struct {
}

func NewSlabClass(chunkSize int) *slabClass {
	return &slabClass{
		ChunkSize: chunkSize,
		slabs:     []*slab{&slab{}},
	}
}

// Allocate new slab.
func (s *slabClass) Grow() {
	s.slabs = append(s.slabs, &slab{})
}
