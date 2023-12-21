package app

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Config Arr

var Logger *logrus.Logger

var DB *gorm.DB
