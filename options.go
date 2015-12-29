package ghalloc

type Options struct {
	// Size of a single slab in bytes.
	// Defaults to 1MB.
	SlabSize int

	// Size of the smallest slab class in bytes.
	// Defaults to 50B.
	MinChunkSize int

	// Step defines Slab classes growing proportion.
	// Must be > 1.0.
	// Defaults to 1.25.
	GrowthFactor float64
}

// Set default values if fields are empty.
func (opt *Options) Init() {
	if opt.SlabSize == 0 {
		opt.SlabSize = 1 * MB
	}

	if opt.MinChunkSize == 0 {
		opt.MinChunkSize = 56 * B
	}

	if opt.GrowthFactor == 0 {
		opt.GrowthFactor = 1.25
	}
}
