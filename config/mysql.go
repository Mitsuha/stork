package config

import "fmt"

var Mysql *mysql

type mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func (m *mysql) ToDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DBName,
	)
}
