package ghalloc

import "testing"

func TestNew(t *testing.T) {
	opt := Options{
		SlabSize:     2 * MB,
		MinChunkSize: 100 * B,
		GrowthFactor: 1.5,
	}

	ghalloc, _ := New(&opt)

	if ghalloc.Opt.SlabSize != 2*MB || ghalloc.Opt.MinChunkSize != 100*B || ghalloc.Opt.GrowthFactor != 1.5 {
		t.Fatalf("new: wrong options: %v", ghalloc.Opt)
	}

	opt.MinChunkSize = 50 * MB
	if _, err := New(&opt); err == nil {
		t.Fatalf("init: expected error")
	}

	opt.MinChunkSize = 0
	opt.GrowthFactor = 0.5
	if _, err := New(&opt); err == nil {
		t.Fatalf("init: expected error")
	}
}

func TestInitSlabClasses(t *testing.T) {
	ghalloc, _ := New(&Options{
		SlabSize:     1 * MB,
		MinChunkSize: 512 * KB,
		GrowthFactor: 2,
	})

	if len(ghalloc.slabClasses) != 2 {
		t.Fatalf("slab classes: wrong number or slub classes: %v", ghalloc.slabClasses)
	}

	if ghalloc.slabClasses[0].ChunkSize != 512*KB && ghalloc.slabClasses[1].ChunkSize != 1*MB {
		t.Fatalf("slab classes: wrong classes")
	}
}