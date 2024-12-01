package usecase

import (
	"context"

	"github.com/d-alejandro/go-code-examples/internal/pkg/dto"
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) GetOrderList(ctx context.Context, req *request.OrderIndexRequest) []*models.Order {
	startValue := req.GetStart()
	limitValue := req.GetEnd() - startValue

	pagination := dto.NewPaginationDTO(
		req.GetSortColumn(),
		req.GetSortType(),
		limitValue,
		startValue,
		req.GetIDs(),
	)

	return useCase.repository.GetOrderList(ctx, pagination)
}
