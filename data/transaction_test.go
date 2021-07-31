package data

import (
	"crypto/sha256"
	"testing"
)

func TestTransaction_Hash(t *testing.T) {
	tests := []struct {
		name string
		tr   TransactionImplSimple
		want Hash
	}{
		{
			"isReturnHash",
			TransactionImplSimple{1111},
			sha256.Sum256([]byte("1111")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.tr
			if got := tr.Hash(); got != tt.want {
				t.Errorf("Transaction.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
