package block

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        int
	Timestamp    int64
	Data         string
	PreviousHash string
	Hash         string
	Nonce        int
}

func NewBlock(index int, data, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
		Hash:         "",
		Nonce:        0,
	}
	block.Hash = block.CalculateHash()
	return block
}

// CalculateHash generates a SHA-256 hash for the block
func (b *Block) CalculateHash() string {
	data := strconv.Itoa(b.Index) + strconv.FormatInt(b.Timestamp, 10) + b.Data + b.PreviousHash + strconv.Itoa(b.Nonce)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// MineBlock implements a simple proof-of-work algorithm
func (b *Block) MineBlock(difficulty int) {
	target := strings.Repeat("0", difficulty)

	fmt.Printf("Mining block %d...\n", b.Index)
	start := time.Now()

	for !strings.HasPrefix(b.Hash, target) {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}

	duration := time.Since(start)
	fmt.Printf("Block %d mined in %v with nonce %d\n", b.Index, duration, b.Nonce)
}

// String returns a formatted string representation of the block
func (b *Block) String() string {
	return fmt.Sprintf("Block #%d\n"+
		"Timestamp: %d\n"+
		"Data: %s\n"+
		"Previous Hash: %s\n"+
		"Hash: %s\n"+
		"Nonce: %d\n",
		b.Index, b.Timestamp, b.Data, b.PreviousHash, b.Hash, b.Nonce)
}
