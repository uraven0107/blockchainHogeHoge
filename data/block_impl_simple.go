package data

import "crypto/sha256"

const transacts_size = 3

type BlockImplSimple struct {
	id           int
	prev         Block
	transactions [transacts_size]Hashable
}

func New(prev Block, t1 Transaction, t2 Transaction, t3 Transaction) BlockImplSimple {
	return BlockImplSimple{
		1,
		prev,
		[transacts_size]Hashable{t1, t2, t3},
	}
}

func (b BlockImplSimple) Hash() Hash {
	var slice []byte
	var pb Block = b.prev
	pb_hash := pb.Hash()
	slice = pb_hash[:]
	for _, t := range b.transactions {
		if t == nil || t.Id() == 0 {
			continue
		}
		t_hash := t.Hash()
		slice = append(slice, t_hash[:]...)
	}
	return sha256.Sum256(slice)
}

func (b BlockImplSimple) Id() int {
	return b.id
}
