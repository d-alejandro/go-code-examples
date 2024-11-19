package providers

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/sirupsen/logrus"
)

const (
	fileFlag       = os.O_RDWR | os.O_CREATE | os.O_APPEND
	filePermission = 0666
)

type LoggerProvider struct {
	cfg *config.Config
}

func NewLoggerProvider(cfg *config.Config) *LoggerProvider {
	return &LoggerProvider{
		cfg: cfg,
	}
}

func (receiver *LoggerProvider) GetLogger() *logrus.Logger {
	logger := logrus.New()

	fileName := fmt.Sprintf("./storage/logs/golang-%s.log", receiver.getCurrentDate(logger))

	file, err := os.OpenFile(fileName, fileFlag, filePermission)
	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}

	if chmodErr := file.Chmod(filePermission); chmodErr != nil {
		logger.Fatalf("error chmod file: %v", chmodErr)
	}

	writer := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(writer)

	receiver.setFormatter(logger)

	return logger
}

func (receiver *LoggerProvider) getCurrentDate(logger *logrus.Logger) string {
	timezone := receiver.cfg.App.TimeZone

	location, err := time.LoadLocation(timezone)
	if err != nil {
		logger.Fatalf("error loading location: %s", err.Error())
	}

	return time.Now().In(location).Format(time.DateOnly)
}

func (*LoggerProvider) setFormatter(logger *logrus.Logger) {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.DateTime

	logger.SetFormatter(formatter)
}
