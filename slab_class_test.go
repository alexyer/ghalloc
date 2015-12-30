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
