package blockchain_user

import (
	"fmt"
	"io"
)

type BlockchainUser struct {
	// blockchain account name
	Name string `json:"name"`
	done bool
}

func NewBlockchainUser() *BlockchainUser {
	return &BlockchainUser{}
}

func (b *BlockchainUser) String() string {
	return fmt.Sprintf("%s", *b)
}

func (bu *BlockchainUser) Read(p []byte) (int, error) {
	if bu.done {
		return 0, io.EOF
	}

	for i, b := range []byte(bu.Name) {
		p[i] = b
	}
	bu.done = true

	return len(bu.Name), nil
}
