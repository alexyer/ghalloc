package ghalloc

type slabClass struct {
	ChunkSize int // Chunk size of the given class.
	SlabSize  int
	slabs     []*slab // Array of slabs of the given class.
}

type slab struct {
	slabClass *slabClass // Slab class of the current slab.
	memory    []byte     // Allocated memory of the given slab.
}

func newSlabClass(chunkSize, slabSize int) *slabClass {
	sc := &slabClass{
		ChunkSize: chunkSize,
		SlabSize:  slabSize,
	}

	sc.slabs = []*slab{newSlab(sc)}

	return sc
}

func newSlab(sc *slabClass) *slab {
	return &slab{
		slabClass: sc,
		memory:    make([]byte, sc.SlabSize),
	}
}

// Allocate new slab.
func (s *slabClass) grow() {
	s.slabs = append(s.slabs, newSlab(s))
}
