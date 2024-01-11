package use_cases

import "github.com/d-alejandro/go-code-examples/internal/app/repositories/interfaces"

type OrderIndexUseCase struct {
	orderIndexRepository interfaces.OrderIndexRepositoryInterface
}

func NewOrderIndexUseCase(orderIndexRepository interfaces.OrderIndexRepositoryInterface) *OrderIndexUseCase {
	return &OrderIndexUseCase{
		orderIndexRepository: orderIndexRepository,
	}
}

func (useCase *OrderIndexUseCase) Execute() string {
	return useCase.orderIndexRepository.Make()
}
