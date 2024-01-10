package countmin

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type countmin struct {
	filterSize     int
	filterItemSize int
	filter         [][]uint32
}

func (c *countmin) generateHash(str string) []uint32 {
	data := []byte(str)
	hashValues := []uint32{}
	for _, v := range SEED {
		hashValues = append(hashValues, murmur3.Sum32WithSeed(data, v))
	}
	return hashValues
}

func minFunc(a uint32, b uint32) uint32 {
	if a > b {
		return b
	}
	return a
}

func NewCountMin(size int) *countmin {
	seedLen := len(SEED)
	filter := make([][]uint32, seedLen)
	for i := 0; i < seedLen; i++ {
		filter[i] = make([]uint32, size)
	}
	return &countmin{
		filterSize:     seedLen,
		filterItemSize: size,
		filter:         filter,
	}
}

func (c *countmin) SetKeyCount(key string) {
	hashValues := c.generateHash(key)
	for i, hashValue := range hashValues {
		index := hashValue % uint32(c.filterItemSize)
		c.filter[i][index] = c.filter[i][index] + 1
	}
}

func (c *countmin) GetKeyCount(key string) uint32 {
	hashValues := c.generateHash(key)
	var min uint32 = math.MaxUint32
	for i, hashValue := range hashValues {
		index := hashValue % uint32(c.filterItemSize)
		min = minFunc(min, c.filter[i][index])
	}
	return min
}
