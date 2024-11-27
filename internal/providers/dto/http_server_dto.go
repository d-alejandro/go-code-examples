package dto

import (
	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/d-alejandro/go-code-examples/internal/providers/bindings"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type HTTPServerDTO struct {
	cfg     *config.Config
	logger  *logrus.Logger
	db      *sqlx.DB
	handler *bindings.HandlerProvider
}

func NewHTTPServerDTO(
	cfg *config.Config,
	logger *logrus.Logger,
	db *sqlx.DB,
	handler *bindings.HandlerProvider,
) *HTTPServerDTO {
	return &HTTPServerDTO{
		cfg:     cfg,
		logger:  logger,
		db:      db,
		handler: handler,
	}
}

func (dto *HTTPServerDTO) GetConfig() *config.Config {
	return dto.cfg
}

func (dto *HTTPServerDTO) GetLogger() *logrus.Logger {
	return dto.logger
}

func (dto *HTTPServerDTO) GetDB() *sqlx.DB {
	return dto.db
}

func (dto *HTTPServerDTO) GetHandler() *bindings.HandlerProvider {
	return dto.handler
}
