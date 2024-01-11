package providers

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
	_ "time/tzdata"
)

const (
	fileFlag       = os.O_RDWR | os.O_CREATE | os.O_APPEND
	filePermission = 0666
	layoutISO      = "2006-01-02"
)

var config *Config
var logger *logrus.Logger

func InitLogger(cfg *Config) *logrus.Logger {
	config = cfg
	logger = logrus.New()

	fileName := getFileName()
	file, err := os.OpenFile(fileName, fileFlag, filePermission)
	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}

	chmodErr := file.Chmod(filePermission)
	if chmodErr != nil {
		logger.Fatalf("error chmod file: %v", chmodErr)
	}

	writer := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(writer)

	return logger
}

func getFileName() string {
	date := getCurrentDate()
	return "./storage/logs/golang-" + date + ".log"
}

func getCurrentDate() string {
	timezone := config.Get("app.timezone").(string)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		logger.Fatal(err.Error())
	}

	timeNow := time.Now()
	return timeNow.In(location).Format(layoutISO)
}
