package usecase

import "github.com/d-alejandro/go-code-examples/internal/pkg/models"

func (useCase *orderUseCase) Delete(id int) (*models.Order, error) {
	response, err := useCase.GetOrder(id)
	if err != nil {
		return response, err
	}

	err = useCase.repository.Delete(response)

	return response, err
}
