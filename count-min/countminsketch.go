package countmin

type countmin struct {
	size   int
	filter [][]uint32
}

func NewCountMin(size int) *countmin {
	filter := make([][]uint32, 3)
	for i := 0; i < 3; i++ {
		filter[i] = make([]uint32, size)
	}
	return &countmin{
		size:   size,
		filter: filter,
	}
}
