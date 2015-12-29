package ghalloc

import "testing"

func TestInit(t *testing.T) {
	opt := Options{}
	opt.Init()

	if opt.SlabSize == 0 || opt.MinChunkSize == 0 || opt.GrowthFactor == 0 {
		t.Fatalf("options: defaults are not set: %v", opt)
	}
}
