package handler

import (
	"os"
	"testing"

	"github.com/d-alejandro/go-code-examples/internal/app/order/presenter"
	"github.com/d-alejandro/go-code-examples/internal/app/order/repository"
	"github.com/d-alejandro/go-code-examples/internal/app/order/usecase"
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/pkg/helpers"
	"github.com/d-alejandro/go-code-examples/internal/providers"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	uCase   usecase.OrderUseCase
	present presenter.OrderPresenter
)

func TestMain(m *testing.M) {
	cfg := config.NewConfig()

	databaseProvider := providers.NewDatabaseProvider(cfg)
	db, err := databaseProvider.GetDB()

	if err != nil {
		logrus.Fatalf("failed to connect SQLX: %s", err.Error())
	}

	rep := repository.NewOrderRepository(db)
	uCase = usecase.NewOrderUseCase(rep)
	rendering := helpers.NewRenderingHelper()
	present = presenter.NewOrderPresenter(rendering)

	exitCode := m.Run()
	os.Exit(exitCode)
}
