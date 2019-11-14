package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Local *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Local: openDB(),
	}
}

func openDB() *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", viper.GetString("db.username"))
	}

	db.DB().SetMaxOpenConns(2000)
	db.DB().SetMaxIdleConns(0)

	return db
}

func (db *Database) Close() {
	DB.Local.Close()
}
