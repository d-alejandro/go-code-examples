package bootstrap

import (
	"github.com/d-alejandro/go-code-examples/internal/app"
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

func InitLogger() *os.File {
	app.Logger = logrus.New()

	fileName := getFileName()
	file, err := os.OpenFile(fileName, fileFlag, filePermission)
	if err != nil {
		app.Logger.Fatalf("error opening file: %v", err)
	}

	chmodErr := file.Chmod(filePermission)
	if chmodErr != nil {
		app.Logger.Fatalf("error chmod file: %v", chmodErr)
	}

	writer := io.MultiWriter(os.Stdout, file)
	app.Logger.SetOutput(writer)

	return file
}

func getFileName() string {
	date := getCurrentDate()
	return "./storage/logs/golang-" + date + ".log"
}

func getCurrentDate() string {
	timezone := app.Config["app"].(app.Arr)["timezone"].(string)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		app.Logger.Fatal(err.Error())
	}

	timeNow := time.Now()
	return timeNow.In(location).Format(layoutISO)
}
