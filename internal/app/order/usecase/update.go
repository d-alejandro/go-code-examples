package usecase

import (
	"github.com/d-alejandro/go-code-examples/internal/pkg/models"
	"github.com/d-alejandro/go-code-examples/internal/pkg/request"
)

func (useCase *orderUseCase) Update(req *request.OrderUpdateRequest, id int) (*models.Order, error) {
	response, err := useCase.GetOrder(id)
	if err != nil {
		return response, err
	}

	err = useCase.repository.Update(req, response)

	return response, err
}
