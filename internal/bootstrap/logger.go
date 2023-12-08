package bootstrap

import (
	"github.com/d-alejandro/go-code-examples/internal/app/identifiers"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func InitLogger() *os.File {
	identifiers.Logger = logrus.New()

	file, err := os.OpenFile("./storage/logs/logrus.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		identifiers.Logger.Fatalf("error opening file: %v", err)
	}

	chmodErr := file.Chmod(0666)
	if chmodErr != nil {
		identifiers.Logger.Fatalf("error chmod file: %v", chmodErr)
	}

	writer := io.MultiWriter(os.Stdout, file)
	identifiers.Logger.SetOutput(writer)

	return file
}
