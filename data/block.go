package data

import "crypto/sha256"

const transacts_size = 3

type Block struct {
	id           int
	transactions [transacts_size]Hashable
}

func New(t1 *Transaction, t2 *Transaction, t3 *Transaction) *Block {
	return &Block{
		1,
		[transacts_size]Hashable{t1, t2, t3},
	}
}

func (b *Block) Hash() [32]byte {
	var slice []byte
	for _, t := range b.transactions {
		if t == nil || t.Id() == 0 {
			continue
		}
		t_hash := t.Hash()
		slice = append(slice, t_hash[:]...)
	}
	return sha256.Sum256(slice)
}

func (b *Block) Id() int {
	return b.id
}
