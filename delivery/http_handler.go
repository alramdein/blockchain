package api

import (
	"net/http"

	"github.com/alramdein/blockchain/usecase"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	Echo              *echo.Echo
	blockchainUsecase *usecase.BlockchainUsecase
}

type TransferRequest struct {
	From   string `json:"from" validate:"required"`
	To     string `json:"to" validate:"required"`
	Amount int    `json:"amount" validate:"required,min=1"`
}

type TransferResponse struct {
	Message    string `json:"message"`
	From       string `json:"from"`
	To         string `json:"to"`
	Amount     int    `json:"amount"`
	BlockIndex int    `json:"block_index"`
}

type ChainResponse struct {
	Chain      []BlockInfo `json:"chain"`
	Length     int         `json:"length"`
	IsValid    bool        `json:"is_valid"`
	Difficulty int         `json:"difficulty"`
}

type BlockInfo struct {
	Index        int    `json:"index"`
	Timestamp    int64  `json:"timestamp"`
	Data         string `json:"data"`
	PreviousHash string `json:"previous_hash"`
	Hash         string `json:"hash"`
	Nonce        int    `json:"nonce"`
}

func NewHTTPHandler(e *echo.Echo) *HTTPHandler {
	return &HTTPHandler{
		Echo:              e,
		blockchainUsecase: usecase.NewBlockchainUsecase(),
	}
}

func (h *HTTPHandler) RegisterRoutes() {
	// Define your HTTP routes and handlers here
	h.Echo.GET("/health", h.Health)
	h.Echo.POST("/transfer", h.Transfer)
	h.Echo.GET("/chain", h.GetChain)
	h.Echo.GET("/balance/:address", h.GetBalance)
}

// Health godoc
// @Summary Health check endpoint
// @Description Returns the health status of the API
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (h *HTTPHandler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "healthy",
		"service": "blockchain-api",
	})
}

// Transfer godoc
// @Summary Transfer coins between addresses
// @Description Create a new transaction to transfer coins from one address to another
// @Tags Blockchain
// @Accept json
// @Produce json
// @Param transfer body TransferRequest true "Transfer details"
// @Success 200 {object} TransferResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transfer [post]
func (h *HTTPHandler) Transfer(c echo.Context) error {
	var req TransferRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	// Validate request
	if req.From == "" || req.To == "" || req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "From, To, and Amount (>0) are required",
		})
	}

	// Execute transfer
	err := h.blockchainUsecase.Transfer(req.From, req.To, req.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Get current chain length to return the block index
	blockchain := h.blockchainUsecase.GetChain()
	blockIndex := len(blockchain.Chain) - 1

	response := TransferResponse{
		Message:    "Transfer completed successfully",
		From:       req.From,
		To:         req.To,
		Amount:     req.Amount,
		BlockIndex: blockIndex,
	}

	return c.JSON(http.StatusOK, response)
}

// GetChain godoc
// @Summary Get the entire blockchain
// @Description Returns the complete blockchain with all blocks and validation status
// @Tags Blockchain
// @Accept json
// @Produce json
// @Success 200 {object} ChainResponse
// @Router /chain [get]
func (h *HTTPHandler) GetChain(c echo.Context) error {
	blockchain := h.blockchainUsecase.GetChain()

	// Convert blockchain to JSON-friendly format
	var blocks []BlockInfo
	for _, block := range blockchain.Chain {
		blockInfo := BlockInfo{
			Index:        block.Index,
			Timestamp:    block.Timestamp,
			Data:         block.Data,
			PreviousHash: block.PreviousHash,
			Hash:         block.Hash,
			Nonce:        block.Nonce,
		}
		blocks = append(blocks, blockInfo)
	}

	response := ChainResponse{
		Chain:      blocks,
		Length:     len(blocks),
		IsValid:    h.blockchainUsecase.ValidateChain(),
		Difficulty: blockchain.Difficulty,
	}

	return c.JSON(http.StatusOK, response)
}

// GetBalance godoc
// @Summary Get balance for an address
// @Description Returns the current balance for the specified address
// @Tags Blockchain
// @Accept json
// @Produce json
// @Param address path string true "Address to check balance for"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /balance/{address} [get]
func (h *HTTPHandler) GetBalance(c echo.Context) error {
	address := c.Param("address")
	if address == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Address parameter is required",
		})
	}

	balance := h.blockchainUsecase.GetBalance(address)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"address": address,
		"balance": balance,
	})
}
