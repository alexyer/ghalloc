package ghalloc

const (
	B      = 1
	KB int = 1 << (10 * iota)
	MB     // Is enough for everyone ;)
)

const (
	CHUNK_ALIGN_BYTES int = 8
)
