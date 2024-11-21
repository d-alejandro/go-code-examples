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
	cfg    *config.Config
	logger *logrus.Logger
}

func NewLoggerProvider(cfg *config.Config) *LoggerProvider {
	return &LoggerProvider{
		cfg:    cfg,
		logger: logrus.New(),
	}
}

func (receiver *LoggerProvider) GetLogger() *logrus.Logger {
	fileName := fmt.Sprintf("./storage/logs/golang-%s.log", receiver.getCurrentDate())

	file, err := os.OpenFile(fileName, fileFlag, filePermission)
	if err != nil {
		receiver.logger.Fatalf("error opening file: %v", err)
	}

	if chmodErr := file.Chmod(filePermission); chmodErr != nil {
		receiver.logger.Fatalf("error chmod file: %v", chmodErr)
	}

	writer := io.MultiWriter(os.Stdout, file)
	receiver.logger.SetOutput(writer)

	receiver.setFormatter()

	return receiver.logger
}

func (receiver *LoggerProvider) getCurrentDate() string {
	location, err := time.LoadLocation(receiver.cfg.App.TimeZone)

	if err != nil {
		receiver.logger.Fatalf("error loading location: %s", err.Error())
	}

	return time.Now().In(location).Format(time.DateOnly)
}

func (receiver *LoggerProvider) setFormatter() {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.DateTime
	receiver.logger.SetFormatter(formatter)
}
