package usecases

import "github.com/d-alejandro/go-code-examples/internal/app/models"

type OrderDestroyUseCase struct {
	showUseCase            OrderShowDestroyUseCaseInterface
	orderDestroyRepository OrderDestroyRepositoryInterface
}

func NewOrderDestroyUseCase(
	showUseCase OrderShowDestroyUseCaseInterface,
	orderDestroyRepository OrderDestroyRepositoryInterface,
) *OrderDestroyUseCase {
	return &OrderDestroyUseCase{
		showUseCase:            showUseCase,
		orderDestroyRepository: orderDestroyRepository,
	}
}

func (useCase *OrderDestroyUseCase) Execute(id int) (*models.Order, error) {
	response, err := useCase.showUseCase.Execute(id)
	if err != nil {
		return response, err
	}

	repositoryError := useCase.orderDestroyRepository.Make(response)

	return response, repositoryError
}
