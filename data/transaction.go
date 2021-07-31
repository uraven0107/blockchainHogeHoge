package data

import (
	"crypto/sha256"
	"strconv"
)

type Transaction struct {
	id int
}

func (t *Transaction) Hash() [32]byte {
	return sha256.Sum256([]byte(strconv.Itoa(t.id)))
}

func (t *Transaction) Id() int {
	return t.id
}
