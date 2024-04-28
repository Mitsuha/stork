package container

import (
	"gorm.io/gorm"
)

var Singled *singleton

type singleton struct {
	DB *gorm.DB
}
