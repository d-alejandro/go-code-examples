package handler

import (
	"os"
	"testing"

	"github.com/d-alejandro/go-code-examples/internal/app/order/repository"
	"github.com/d-alejandro/go-code-examples/internal/app/order/usecase"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/providers"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	repos repository.OrderRepository
	uCase usecase.OrderUseCase
)

func TestMain(m *testing.M) {
	cfg := config.NewConfig()

	databaseProvider := providers.NewDatabaseProvider(cfg)
	db, err := databaseProvider.GetDB()

	if err != nil {
		logrus.Fatalf("failed to connect SQLX: %s", err.Error())
	}

	repos = repository.NewOrderRepository(db)
	uCase = usecase.NewOrderUseCase(repos)

	exitCode := m.Run()
	os.Exit(exitCode)
}
