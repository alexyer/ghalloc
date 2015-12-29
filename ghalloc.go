// Implementation of Memcached inspired Slab Allocator.
package ghalloc

import "errors"

type Ghalloc struct {
	Opt         *Options
	slabClasses []*slabClass // Array of slabClasses. Each slabClass is an array of slabs.
}

func New(opt *Options) (*Ghalloc, error) {
	opt.Init()

	if opt.SlabSize < opt.MinChunkSize {
		return nil, errors.New("ghalloc: minimum chunk must be less than slab size")
	}

	if opt.GrowthFactor < 1 {
		return nil, errors.New("ghalloc: growth factor must be > 1")
	}

	g := &Ghalloc{
		Opt: opt,
	}

	g.InitSlabsClasses()

	return g, nil
}

// Create and init slab classes for each chunk size.
func (g *Ghalloc) InitSlabsClasses() {
	var size int
	for size = g.Opt.MinChunkSize; size <= int(float64(g.Opt.SlabSize)/g.Opt.GrowthFactor); {
		// Make sure items are always n-byte aligned
		if size%CHUNK_ALIGN_BYTES != 0 {
			size += CHUNK_ALIGN_BYTES - (size % CHUNK_ALIGN_BYTES)
		}

		g.slabClasses = append(g.slabClasses, NewSlabClass(size, g.Opt.SlabSize))

		size = int(float64(size) * g.Opt.GrowthFactor)
	}

	g.slabClasses = append(g.slabClasses, NewSlabClass(g.Opt.SlabSize, g.Opt.SlabSize))
}
