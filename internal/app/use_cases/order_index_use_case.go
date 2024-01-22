package use_cases

import (
	"github.com/d-alejandro/go-code-examples/internal/app/dto"
	"github.com/d-alejandro/go-code-examples/internal/app/http/requests"
	"github.com/d-alejandro/go-code-examples/internal/app/models"
	"github.com/d-alejandro/go-code-examples/internal/app/repositories"
)

type OrderIndexRepository interface {
	Make(pagination repositories.PaginationDTO) []models.Order
}

type OrderIndexUseCase struct {
	orderIndexRepository OrderIndexRepository
}

func NewOrderIndexUseCase(repository OrderIndexRepository) *OrderIndexUseCase {
	return &OrderIndexUseCase{repository}
}

func (useCase *OrderIndexUseCase) Execute(request requests.OrderIndexRequest) []models.Order {
	limitValue := request.End - request.Start

	pagination := dto.NewPaginationDTO(
		request.SortColumn,
		request.SortType,
		limitValue,
		request.Start,
	)

	return useCase.orderIndexRepository.Make(pagination)
}
