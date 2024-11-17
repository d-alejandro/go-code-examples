package usecase

import (
	"github.com/d-alejandro/go-code-examples/internal/app/order/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
)

type OrderIndexUseCase struct {
	orderIndexRepository OrderIndexRepositoryInterface
}

func NewOrderIndexUseCase(repository OrderIndexRepositoryInterface) *OrderIndexUseCase {
	return &OrderIndexUseCase{repository}
}

func (useCase *OrderIndexUseCase) Execute(request interface{ OrderIndexRequestInterface }) []*models.Order {
	startValue := request.GetStart()
	limitValue := request.GetEnd() - startValue

	pagination := dto.NewPaginationDTO(
		request.GetSortColumn(),
		request.GetSortType(),
		limitValue,
		startValue,
	)

	return useCase.orderIndexRepository.Make(pagination)
}
