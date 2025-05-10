package packcalculator

import (
	"reflect"
	"testing"
)

// TestCalculateOptimalPacks runs different scenarios to make sure the pack calculator
// returns the right combination with the smallest total and pack count.
func TestCalculateOptimalPacks(t *testing.T) {
	calculator := NewCalculatorService()

	tests := []struct {
		name          string
		items         int
		packSizes     []int
		expectedPacks map[int]int
		expectedTotal int
		expectedCount int
	}{
		{
			name:          "Exact match single pack",
			items:         250,
			packSizes:     []int{250, 500, 1000},
			expectedPacks: map[int]int{250: 1},
			expectedTotal: 250,
			expectedCount: 1,
		},
		{
			name:          "Needs smallest overage",
			items:         251,
			packSizes:     []int{250, 500},
			expectedPacks: map[int]int{500: 1},
			expectedTotal: 500,
			expectedCount: 1,
		},
		{
			name:          "Multiple packs - optimal",
			items:         1201,
			packSizes:     []int{250, 500, 1000, 2000},
			expectedPacks: map[int]int{1000: 1, 250: 1},
			expectedTotal: 1250,
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculator.CalculateOptimalPacks(tt.items, tt.packSizes)

			if !reflect.DeepEqual(result.Packs, tt.expectedPacks) {
				t.Errorf("Expected packs %v, got %v", tt.expectedPacks, result.Packs)
			}
			if result.TotalItems != tt.expectedTotal {
				t.Errorf("Expected total items %d, got %d", tt.expectedTotal, result.TotalItems)
			}
			if result.TotalPacks != tt.expectedCount {
				t.Errorf("Expected total packs %d, got %d", tt.expectedCount, result.TotalPacks)
			}
		})
	}
}
