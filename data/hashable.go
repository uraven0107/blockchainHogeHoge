package data

type Hashable interface {
	Hash() Hash
	Id() int
}
