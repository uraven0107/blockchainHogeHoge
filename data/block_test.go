package data

import (
	"crypto/sha256"
	"reflect"
	"strconv"
	"testing"
)

func TestBlock_Hash(t *testing.T) {
	const id1 = 123
	const id2 = 456
	const id3 = 789
	var id_hash1 [32]byte = sha256.Sum256([]byte(strconv.Itoa(id1)))
	var id_hash2 [32]byte = sha256.Sum256([]byte(strconv.Itoa(id2)))
	var id_hash3 [32]byte = sha256.Sum256([]byte(strconv.Itoa(id3)))
	tests := []struct {
		name      string
		transacts [3]Transaction
		want      [32]byte
	}{
		{
			"isReturnHashCaseOfOneTransaction",
			[3]Transaction{
				{id1},
			},
			sha256.Sum256((id_hash1[:])),
		},
		{
			"isReturnHashCaseOfThreeTransaction",
			[3]Transaction{
				{id1},
				{id2},
				{id3},
			},
			sha256.Sum256((append(append(id_hash1[:], id_hash2[:]...), id_hash3[:]...))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New(&tt.transacts[0], &tt.transacts[1], &tt.transacts[2])
			if got := b.Hash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Block.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
