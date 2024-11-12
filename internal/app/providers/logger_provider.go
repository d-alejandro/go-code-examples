package providers

import (
	"io"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/d-alejandro/go-code-examples/internal/config"
	"github.com/sirupsen/logrus"
)

const (
	fileFlag       = os.O_RDWR | os.O_CREATE | os.O_APPEND
	filePermission = 0666
)

type LoggerProvider struct {
	config *config.Config
	logger *logrus.Logger
}

func NewLoggerProvider(config *config.Config) *LoggerProvider {
	loggerProvider := &LoggerProvider{
		config: config,
	}
	loggerProvider.register()
	return loggerProvider
}

func (receiver *LoggerProvider) GetLogger() *logrus.Logger {
	return receiver.logger
}

func (receiver *LoggerProvider) register() {
	receiver.logger = logrus.New()

	fileName := receiver.getFileName()
	file, err := os.OpenFile(fileName, fileFlag, filePermission)
	if err != nil {
		receiver.logger.Fatalf("error opening file: %v", err)
	}

	chmodErr := file.Chmod(filePermission)
	if chmodErr != nil {
		receiver.logger.Fatalf("error chmod file: %v", chmodErr)
	}

	writer := io.MultiWriter(os.Stdout, file)
	receiver.logger.SetOutput(writer)

	receiver.setFormatter()
}

func (receiver *LoggerProvider) getFileName() string {
	date := receiver.getCurrentDate()
	return "./storage/logs/golang-" + date + ".log"
}

func (receiver *LoggerProvider) getCurrentDate() string {
	timezone := receiver.config.App.TimeZone

	location, err := time.LoadLocation(timezone)
	if err != nil {
		receiver.logger.Fatal(err.Error())
	}

	timeNow := time.Now()
	return timeNow.In(location).Format(time.DateOnly)
}

func (receiver *LoggerProvider) setFormatter() {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.DateTime
	receiver.logger.SetFormatter(formatter)
}
