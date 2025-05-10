package packcalculator

type calculatorService interface {
	CalculateOptimalPacks(items int, packSizes []int) PackResult
}

// UseCase is struct is our use case – it connects the business logic (what we want to do)
// with the actual calculation logic (how it gets done).
type UseCase struct {
	calculator calculatorService
}

// NewUseCase is just sets up a new instance of the use case.
// We pass in the calculator dependency here (typical dependency injection).
func NewUseCase(calculator *CalculatorService) *UseCase {
	return &UseCase{
		calculator: calculator,
	}
}

// Calculate is the main method for the use case.
// It delegates the calculation to the service layer.
func (uc *UseCase) Calculate(items int, packSizes []int) PackResult {
	return uc.calculator.CalculateOptimalPacks(items, packSizes)
}
