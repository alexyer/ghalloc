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
	for i := 0; i < sc.Capacity; i++ {
		s.chunks[i] = unsafe.Pointer(uintptr(unsafe.Pointer(&s.memory[0])) + unsafe.Sizeof(s.memory[0])*uintptr(i))
	}

	return s
}
