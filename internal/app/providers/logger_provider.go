package providers

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
	_ "time/tzdata"
)

type LoggerProvider struct {
	configProvider *ConfigProvider
	logger         *logrus.Logger
}

const (
	fileFlag       = os.O_RDWR | os.O_CREATE | os.O_APPEND
	filePermission = 0666
	layoutISO      = "2006-01-02"
)

func NewLoggerProvider(configProvider *ConfigProvider) *LoggerProvider {
	loggerProvider := &LoggerProvider{
		configProvider: configProvider,
	}
	loggerProvider.register()
	return loggerProvider
}

func GetLogger(loggerProvider *LoggerProvider) *logrus.Logger {
	return loggerProvider.logger
}

func (loggerProvider *LoggerProvider) register() {
	loggerProvider.logger = logrus.New()

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
}

func (loggerProvider *LoggerProvider) getFileName() string {
	date := loggerProvider.getCurrentDate()
	return "./storage/logs/golang-" + date + ".log"
}

func (loggerProvider *LoggerProvider) getCurrentDate() string {
	timezone := loggerProvider.configProvider.GetConfig("app.timezone").(string)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		loggerProvider.logger.Fatal(err.Error())
	}

	timeNow := time.Now()
	return timeNow.In(location).Format(layoutISO)
}
