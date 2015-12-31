package ghalloc

import "testing"

func TestGrow(t *testing.T) {
	sc := newSlabClass(80, 1*MB)
	slabsLen := len(sc.slabs)
	sc.grow()

	if len(sc.slabs) != slabsLen+1 {
		t.Fatal("slab class: grow: error")
	}
}

func TestFindAvailableSlab(t *testing.T) {
	sc := newSlabClass(512*KB, 1*MB)
	sc.grow()
	sc.slabs[0].full = true

	slab := sc.findAvailableSlab()

	if slab != sc.slabs[1] {
		t.Fatal("slab class: find slab: found wrong slab")
	}

	for _, slab := range sc.slabs {
		slab.full = true
	}

	slab = sc.findAvailableSlab()

	if slab != nil {
		t.Fatal("slab class: find slab: found wrong slab. Expected nil")
	}
}
