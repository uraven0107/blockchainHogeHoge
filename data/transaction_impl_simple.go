package data

import (
	"crypto/sha256"
	"strconv"
)

type TransactionImplSimple struct {
	id int
}

func (t TransactionImplSimple) Hash() Hash {
	return sha256.Sum256([]byte(strconv.Itoa(t.id)))
}

func (t TransactionImplSimple) Id() int {
	return t.id
}
