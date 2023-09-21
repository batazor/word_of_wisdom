package proofofwork

import (
	"github.com/batazor/word_of_wisdow/internal/domain/block"
	"github.com/batazor/word_of_wisdow/internal/domain/proofofwork/simple"
)

// New returns a new PoW instance.
func New(b *block.Block) (PoW, error) {
	pow, err := simple.New(b)
	if err != nil {
		return nil, err
	}

	return pow, nil
}
