package bootstrap

import (
	"github.com/d-alejandro/go-code-examples/internal/app/providers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

var (
	connection sync.Once
	config     *providers.Config
	logger     *logrus.Logger
	database   *gorm.DB
)

func Boot() (*providers.Config, *logrus.Logger, *gorm.DB) {
	connection.Do(func() {
		config = providers.NewConfig()
		logger = providers.InitLogger(config)
		database = providers.GetDatabase(config, logger)
	})
	return config, logger, database
}
