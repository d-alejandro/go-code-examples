package providers

import (
	"github.com/d-alejandro/go-code-examples/internal/app/helpers"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
	_ "time/tzdata"
)

type LoggerProvider struct {
	config *helpers.Config
	logger *logrus.Logger
}

func NewLoggerProvider(container *helpers.DependenciesContainer) *LoggerProvider {
	config := container.GetDependency("config").(*helpers.Config)
	loggerProvider := &LoggerProvider{
		config: config,
	}
	loggerProvider.register()
	return loggerProvider
}

func (loggerProvider *LoggerProvider) GetLogger() *logrus.Logger {
	return loggerProvider.logger
}

func (loggerProvider *LoggerProvider) register() {
	loggerProvider.logger = logrus.New()

	const (
		fileFlag       = os.O_RDWR | os.O_CREATE | os.O_APPEND
		filePermission = 0666
	)

	fileName := loggerProvider.getFileName()
	file, err := os.OpenFile(fileName, fileFlag, filePermission)
	if err != nil {
		loggerProvider.logger.Fatalf("error opening file: %v", err)
	}

	chmodErr := file.Chmod(filePermission)
	if chmodErr != nil {
		loggerProvider.logger.Fatalf("error chmod file: %v", chmodErr)
	}

	writer := io.MultiWriter(os.Stdout, file)
	loggerProvider.logger.SetOutput(writer)

	loggerProvider.setFormatter()
}

func (loggerProvider *LoggerProvider) getFileName() string {
	date := loggerProvider.getCurrentDate()
	return "./storage/logs/golang-" + date + ".log"
}

func (loggerProvider *LoggerProvider) getCurrentDate() string {
	timezone := loggerProvider.config.Get("app.timezone").(string)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		loggerProvider.logger.Fatal(err.Error())
	}

	timeNow := time.Now()
	return timeNow.In(location).Format(time.DateOnly)
}

func (loggerProvider *LoggerProvider) setFormatter() {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.DateTime
	loggerProvider.logger.SetFormatter(formatter)
}
