package data

import (
	"crypto/sha256"
	"reflect"
	"strconv"
	"testing"
)

type block_mock struct {
	Block
}

var hogehoge_hash = sha256.Sum256([]byte("hogehoge"))

func (b block_mock) Hash() Hash {
	return hogehoge_hash
}

func TestBlock_Hash(t *testing.T) {
	const id1 = 123
	const id2 = 456
	const id3 = 789
	var id_hash1 Hash = sha256.Sum256([]byte(strconv.Itoa(id1)))
	var id_hash2 Hash = sha256.Sum256([]byte(strconv.Itoa(id2)))
	var id_hash3 Hash = sha256.Sum256([]byte(strconv.Itoa(id3)))
	tests := []struct {
		name      string
		transacts [3]Transaction
		want      Hash
	}{
		{
			"isReturnHashCaseOfOneTransaction",
			[3]Transaction{
				TransactionImplSimple{id1},
			},
			sha256.Sum256(append(hogehoge_hash[:], id_hash1[:]...)),
		},
		{
			"isReturnHashCaseOfThreeTransaction",
			[3]Transaction{
				TransactionImplSimple{id1},
				TransactionImplSimple{id2},
				TransactionImplSimple{id3},
			},
			sha256.Sum256((append(append(append(hogehoge_hash[:], id_hash1[:]...), id_hash2[:]...), id_hash3[:]...))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := new(block_mock)
			b := New(prev, tt.transacts[0], tt.transacts[1], tt.transacts[2])
			if got := b.Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Block.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
