package bitarray

type BitArray interface {
}

type bitArray struct {
	length uint
	blocks []uint64
}
