// Implementation of Memcached inspired Slab Allocator.
package ghalloc

import (
	"errors"
	"fmt"
	"sync"
	"unsafe"
)

type Ghalloc struct {
	Opt              *Options
	slabClassesLocks []sync.Mutex
	slabClasses      []*slabClass // Array of slabClasses. Each slabClass is an array of slabs.
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

		g.slabClasses = append(g.slabClasses, newSlabClass(size, g.Opt.SlabSize))

		size = int(float64(size) * g.Opt.GrowthFactor)
	}

	g.slabClasses = append(g.slabClasses, newSlabClass(g.Opt.SlabSize, g.Opt.SlabSize))
	g.slabClassesLocks = make([]sync.Mutex, len(g.slabClasses))
}

// Print information about slab classes.
// Debug function.
func (g *Ghalloc) PrintSlabClassesInfo() {
	for _, sc := range g.slabClasses {
		fmt.Printf("Chunk size: %9d, Capacity: %9d\n", sc.ChunkSize, sc.SlabSize/sc.ChunkSize)
	}
}

// Allocate region of the given size.
func (g *Ghalloc) Alloc(size uintptr) unsafe.Pointer {
	if int(size) > g.Opt.SlabSize {
		return nil
	}

	i := g.findSlabClassIndex(int(size))

	// TODO(alexyer): The lock locks the whole slab class.
	//                Think about better granular lock.
	g.slabClassesLocks[i].Lock()
	ptr := g.slabClasses[i].getChunk()
	g.slabClassesLocks[i].Unlock()

	return ptr
}

// Find suitable slab class for the given size.
func (g *Ghalloc) findSlabClassIndex(size int) int {
	for i := 0; i < len(g.slabClasses); i++ {
		if size <= g.slabClasses[i].ChunkSize {
			return i
		}
	}
	return -1
}
