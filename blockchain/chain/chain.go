package chain

import (
	"fmt"

	"github.com/alramdein/blockchain/blockchain/block"
)

type Blockchain struct {
	Chain      []*block.Block
	Difficulty int
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain() *Blockchain {
	bc := &Blockchain{
		Chain:      []*block.Block{},
		Difficulty: 2, // Number of leading zeros required in hash
	}
	bc.createGenesisBlock()
	return bc
}

// createGenesisBlock creates the first block in the blockchain
func (bc *Blockchain) createGenesisBlock() {
	genesis := block.NewBlock(0, "Genesis Block", "0")
	genesis.MineBlock(bc.Difficulty)
	bc.Chain = append(bc.Chain, genesis)
}

// GetLatestBlock returns the most recent block in the chain
func (bc *Blockchain) GetLatestBlock() *block.Block {
	return bc.Chain[len(bc.Chain)-1]
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.GetLatestBlock()
	newBlock := block.NewBlock(
		previousBlock.Index+1,
		data,
		previousBlock.Hash,
	)
	newBlock.MineBlock(bc.Difficulty)
	bc.Chain = append(bc.Chain, newBlock)
}

// IsChainValid validates the entire blockchain
func (bc *Blockchain) IsChainValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		// Check if current block's hash is valid
		if currentBlock.Hash != currentBlock.CalculateHash() {
			fmt.Printf("Invalid hash at block %d\n", i)
			return false
		}

		// Check if current block points to previous block
		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("Invalid previous hash at block %d\n", i)
			return false
		}
	}
	return true
}

// GetBalance calculates balance for a given address (simplified)
func (bc *Blockchain) GetBalance(address string) int {
	balance := 0

	for _, block := range bc.Chain {
		// Simple transaction parsing (in a real blockchain, this would be more complex)
		if block.Data == fmt.Sprintf("Transfer to %s: 10 coins", address) {
			balance += 10
		}
		if block.Data == fmt.Sprintf("Transfer from %s: 10 coins", address) {
			balance -= 10
		}
	}

	return balance
}

// PrintChain displays the entire blockchain
func (bc *Blockchain) PrintChain() {
	for i, block := range bc.Chain {
		fmt.Printf("=== Block %d ===\n", i)
		fmt.Println(block.String())
		fmt.Println()
	}
}

// GetChainLength returns the number of blocks in the chain
func (bc *Blockchain) GetChainLength() int {
	return len(bc.Chain)
}
