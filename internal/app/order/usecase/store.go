package usecase

import "github.com/d-alejandro/go-code-examples/internal/app/order/models"

type OrderStoreUseCase struct {
	orderStoreRepository OrderStoreRepositoryInterface
}

func NewOrderStoreUseCase(orderStoreRepository OrderStoreRepositoryInterface) *OrderStoreUseCase {
	return &OrderStoreUseCase{orderStoreRepository}
}

func (useCase *OrderStoreUseCase) Execute(request interface{ OrderStoreRequestInterface }) (*models.Order, error) {
	return useCase.orderStoreRepository.Make(request)
}
