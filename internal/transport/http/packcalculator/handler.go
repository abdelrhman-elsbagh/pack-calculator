package packcalculator

import (
	"github.com/abdelrhman-elsbagh/pack-calculator/internal/domain/packcalculator"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PackResult represents the final result of the pack calculation.
// Packs: a map of pack size quantity
// TotalItems: total number of items that will be shipped
// TotalPacks: total number of packs used
type PackResult struct {
	Packs      map[int]int `json:"packs"`
	TotalItems int         `json:"total_items"`
	TotalPacks int         `json:"total_packs"`
}

// OrderRequest represents the input we receive from the client/API call.
// It contains the number of items requested and the available pack sizes to use.
type OrderRequest struct {
	Items     int   `json:"items"`      // How many items the user wants
	PackSizes []int `json:"pack_sizes"` // Available pack sizes to choose from
}

// Interface to make testing/decoupling easier.
// It defines what the use case must be able to do.
type packCalculator interface {
	Calculate(items int, packSizes []int) packcalculator.PackResult
}

// PackHandler is handler struct that connects HTTP layer to the business logic.
type PackHandler struct {
	useCase packCalculator
}

// NewPackHandler Constructor for PackHandler. We pass the use case into it here.
func NewPackHandler(useCase packCalculator) *PackHandler {
	return &PackHandler{
		useCase: useCase,
	}
}

// Calculate handles POST /calculate
// It reads the input, validates it, sends it to the use case, and returns the result.
func (h *PackHandler) Calculate(c *gin.Context) {
	var req OrderRequest

	// Try to parse the incoming JSON into our request struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Make sure we actually got valid input (items > 0 and at least one pack size)
	if req.Items <= 0 || len(req.PackSizes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Items and pack_sizes are required"})
		return
	}

	// Pass it to the use case to do the actual logic
	result := h.useCase.Calculate(req.Items, req.PackSizes)
	httpResp := PackResult{
		Packs:      result.Packs,
		TotalItems: result.TotalItems,
		TotalPacks: result.TotalPacks,
	}

	// Send the result back to the client as JSON
	c.JSON(http.StatusOK, httpResp)
}
