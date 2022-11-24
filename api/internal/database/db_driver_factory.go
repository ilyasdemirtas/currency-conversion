package database

import (
	"arf/currency-conversion/api/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conf config.DB

func DbDriver() gorm.Dialector {
	conf = config.NewDB()

	var driver = conf.Driver
	switch driver {
	case "sqlite":
		return sqliteD()
	case "postgresql":
		return postgresqlD()
	}
	return mysqlD()
}

func mysqlD() gorm.Dialector {
	dsn := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.Database +
		"?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=" + conf.AllowNativePasswords
	return mysql.Open(dsn)
}

func postgresqlD() gorm.Dialector {
	dsn := "host=" + conf.Host + " user=" + conf.User + " password=" + conf.Password + " dbname=" +
		conf.Database + " port=" + conf.Port + " sslmode=disable TimeZone=Europe/Istanbul"
	return postgres.Open(dsn)
}

func sqliteD() gorm.Dialector {
	return sqlite.Open(conf.Database)
}
