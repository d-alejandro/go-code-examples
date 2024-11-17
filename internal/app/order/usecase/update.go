package usecase

import "github.com/d-alejandro/go-code-examples/internal/pkg/models"

type OrderUpdateUseCase struct {
	showUseCase           OrderShowUseCaseInterface
	orderUpdateRepository OrderUpdateRepositoryInterface
}

func NewOrderUpdateUseCase(
	showUseCase OrderShowUseCaseInterface,
	orderUpdateRepository OrderUpdateRepositoryInterface,
) *OrderUpdateUseCase {
	return &OrderUpdateUseCase{
		showUseCase:           showUseCase,
		orderUpdateRepository: orderUpdateRepository,
	}
}

func (useCase *OrderUpdateUseCase) Execute(
	request interface{ OrderUpdateRequestInterface },
	id int,
) (*models.Order, error) {
	response, err := useCase.showUseCase.Execute(id)
	if err != nil {
		return response, err
	}

	repositoryError := useCase.orderUpdateRepository.Make(request, response)

	return response, repositoryError
}
