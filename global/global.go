package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sso/config"
)

var (
	CONFIG config.Serve
	LOG    *logrus.Logger
	DB     *gorm.DB
)
