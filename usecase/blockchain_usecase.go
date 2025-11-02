package usecase

import (
	"fmt"

	"github.com/alramdein/blockchain/blockchain/chain"
)

type BlockchainUsecase struct {
	blockchain *chain.Blockchain
}

func NewBlockchainUsecase() *BlockchainUsecase {
	return &BlockchainUsecase{
		blockchain: chain.NewBlockchain(),
	}
}

func (uc *BlockchainUsecase) Transfer(from, to string, amount int) error {
	// In a real blockchain, you'd check balances and validate signatures
	// For this simple implementation, we'll just add the transaction
	transferData := fmt.Sprintf("Transfer from %s to %s: %d coins", from, to, amount)
	uc.blockchain.AddBlock(transferData)
	return nil
}

func (uc *BlockchainUsecase) GetChain() *chain.Blockchain {
	return uc.blockchain
}

func (uc *BlockchainUsecase) GetBalance(address string) int {
	return uc.blockchain.GetBalance(address)
}

func (uc *BlockchainUsecase) ValidateChain() bool {
	return uc.blockchain.IsChainValid()
}
