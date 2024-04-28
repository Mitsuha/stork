package container

import (
	"github.com/mitsuha/stork/config"
	"github.com/mitsuha/stork/repository/model/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Boot() error {
	db, err := gorm.Open(mysql.Open(config.Mysql.ToDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	Singled = &singleton{
		DB: db,
	}

	dao.SetDefault(db)

	return nil
}
