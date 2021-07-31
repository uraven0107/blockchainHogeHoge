package data

type Hashable interface {
	Hash() [32]byte
	Id() int
}
