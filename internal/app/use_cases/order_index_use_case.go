package use_cases

type OrderIndexUseCase struct {
}

func NewOrderIndexUseCase() *OrderIndexUseCase {
	return new(OrderIndexUseCase)
}

func (useCase *OrderIndexUseCase) Execute() string {
	return "OrderIndexController run."
}
