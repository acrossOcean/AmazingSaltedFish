package service

import (
	"fmt"
	"sync"

	"github.com/acrossOcean/config"
	"github.com/acrossOcean/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	UserName    string
	Password    string
	DBUrl       string
	DBName      string
	ExtraConfig string
}

var (
	_DB     *gorm.DB
	_initDB sync.Once
)

func InitDB() (err error) {
	_initDB.Do(func() {
		err = InitDB()
	})

	return err
}

func initDB() (err error) {
	connStr := getDBConfig().getConnStr()
	log.Debug("连接数据库:%s", connStr)
	if _DB, err = gorm.Open("mysql", connStr); err != nil {
		log.Error("打开数据库连接错误:%s", err.Error())
		return err
	}

	_DB.DB().SetMaxIdleConns(config.DefaultInt("mysql>>maxIdleConn", 10))
	_DB.DB().SetMaxOpenConns(config.DefaultInt("mysql>>maxOpenConn", 10))
	_DB.LogMode(config.DefaultBool("mysql>>logMode", true))

	return err
}

func GetDB() *gorm.DB {
	_initDB.Do(func() {
		if _DB == nil {
			err := initDB()
			if err != nil {
				log.Error("打开数据库连接错误:%s", err.Error())
			}
		}
	})

	return _DB
}

func CloseDB() error {
	return _DB.Close()
}

func getDBConfig() DBConfig {
	result := DBConfig{}

	result.UserName, _ = config.String("mysql>>username")
	result.Password, _ = config.String("mysql>>password")
	result.DBUrl, _ = config.String("mysql>>dbUrl")
	result.DBName, _ = config.String("mysql>>dbName")
	result.ExtraConfig, _ = config.String("mysql>>extraConfig")

	return result
}

func (receiver DBConfig) getConnStr() string {
	str := fmt.Sprintf("%s:%s@(%s)/%s?%s", receiver.UserName, receiver.Password, receiver.DBUrl, receiver.DBName, receiver.ExtraConfig)
	return str
}
