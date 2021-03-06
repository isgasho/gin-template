package utils

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"

	// _ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", "./gateway.db")
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("初始化数据库")
}
