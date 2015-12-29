package ghalloc

import "testing"

func TestNewSlab(t *testing.T) {
	sc := newSlabClass(42, 1*MB)
	slab := newSlab(sc)

	if len(slab.memory) != 42 {
		t.Fatalf("slab: allocated wrong memory slice")
	}
}

func TestGrow(t *testing.T) {
	sc := newSlabClass(80, 1*MB)
	slabsLen := len(sc.slabs)
	sc.grow()

	if len(sc.slabs) != slabsLen+1 {
		t.Fatalf("slab class: grow: error")
	}
}
