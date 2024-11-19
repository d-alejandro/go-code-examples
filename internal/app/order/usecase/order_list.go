package usecase

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) GetOrderList(request *request.OrderIndexRequest) []*models.Order {
	startValue := request.GetStart()
	limitValue := request.GetEnd() - startValue

	pagination := dto.NewPaginationDTO(
		request.GetSortColumn(),
		request.GetSortType(),
		limitValue,
		startValue,
	)

	return useCase.repository.GetOrderList(pagination)
}
