package packcalculator

import (
	"math"
	"sort"
)

// PackResult represents the final result of the pack calculation.
// Packs: a map of pack size quantity
// TotalItems: total number of items that will be shipped
// TotalPacks: total number of packs used
type PackResult struct {
	Packs      map[int]int
	TotalItems int
	TotalPacks int
}

// CalculatorService is responsible for executing the core logic for calculating optimal pack combinations.
type CalculatorService struct{}

// NewCalculatorService initializes and returns a new instance of CalculatorService.
func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

// CalculateOptimalPacks finds the best combination of packs to fulfill the requested number of items.
// It minimizes total items (overage) and the number of packs used.
func (cs *CalculatorService) CalculateOptimalPacks(items int, packSizes []int) PackResult {
	// Sort pack sizes in descending order to try larger packs first (greedy first).
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	minTotal := math.MaxInt   // Tracks the lowest total items used across all combinations
	minPacks := math.MaxInt   // Tracks the lowest number of packs used
	var bestResult PackResult // Stores the best result found

	// Start recursive exploration of all combinations
	cs.findCombination(items, packSizes, 0, map[int]int{}, 0, 0, &minTotal, &minPacks, &bestResult)

	return bestResult
}

// cloneMap creates a new copy of the given map[int]int.
// This is required to avoid modifying shared memory in recursive calls.
func cloneMap(src map[int]int) map[int]int {
	copy := make(map[int]int)
	for k, v := range src {
		copy[k] = v
	}
	return copy
}

// isBetterSolution determines if the new solution is better than the current best.
// It prefers solutions with fewer total items, and then fewer packs if totals are equal.
func isBetterSolution(newTotal, newPacks, minTotal, minPacks int) bool {
	return newTotal < minTotal || (newTotal == minTotal && newPacks < minPacks)
}

// saveBestResult updates the bestResult with the current best combination found.
func saveBestResult(result *PackResult, current map[int]int, newTotal int, minTotal, minPacks *int) {
	resultCopy := make(map[int]int)
	totalPacks := 0

	for k, v := range current {
		if v > 0 {
			resultCopy[k] = v
			totalPacks += v
		}
	}

	*minTotal = newTotal
	*minPacks = totalPacks
	result.Packs = resultCopy
	result.TotalItems = newTotal
	result.TotalPacks = totalPacks
}

// findCombination is a recursive function that explores all possible combinations of packs.
// It attempts to build up from the current index and tracks the best solution found.
func (cs *CalculatorService) findCombination(
	target int, // Number of items needed
	sizes []int, // Available pack sizes
	index int, // Current index in sizes slice
	current map[int]int, // Current combination of pack size quantity
	currentTotal int, // Current total of items selected
	currentPacks int, // Current total of packs selected
	minTotal *int, // Pointer to best total items (minimized)
	minPacks *int, // Pointer to best number of packs (minimized)
	bestResult *PackResult, // Pointer to the current best result
) {
	// Base case: if we exceeded the index, stop recursion
	if index >= len(sizes) {
		return
	}

	size := sizes[index]
	// Maximum number of packs of this size we can consider without exceeding the target by too much
	maxCount := (target - currentTotal + size - 1) / size

	for count := 0; count <= maxCount; count++ {
		// Create a new copy of the current map to avoid mutation
		newCurrent := cloneMap(current)

		// Add or remove the current pack size depending on the count
		if count > 0 {
			newCurrent[size] = count
		} else {
			delete(newCurrent, size)
		}

		newTotal := currentTotal + count*size
		newPacks := currentPacks + count

		// If we've reached or exceeded the target, check if this is a better solution
		if newTotal >= target {
			if isBetterSolution(newTotal, newPacks, *minTotal, *minPacks) {
				saveBestResult(bestResult, newCurrent, newTotal, minTotal, minPacks)
			}
			continue
		}

		// Recurse to try next pack size
		cs.findCombination(
			target,
			sizes,
			index+1,
			newCurrent,
			newTotal,
			newPacks,
			minTotal,
			minPacks,
			bestResult,
		)
	}
}
