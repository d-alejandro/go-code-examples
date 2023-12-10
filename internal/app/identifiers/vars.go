package identifiers

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Config ConfigMap

var Logger *logrus.Logger

var DBConnection *gorm.DB
