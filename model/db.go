/**
* @file   : init.go
* @descrip: 初始化数据库连接，多数据库连接管理; 总结就是要拿到数据库连接对象
* @author : ch-yk
* @create : 2018-09-02 下午1:32
* @email  : commonheart.yk@gmail.com
**/

package model

import (
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

/*ke'ne可能要应对多个数据库连接，目前就给两个连接实例*/
type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
	//其他同类型(比如关系型)备份数据库
}

var DB *Database

/*内部方法，被 Init 或者 Get类方法调用*/
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	/* 调用 gorm 框架来获取连接实例, mysql驱动适用于 mariadb */
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(1000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(10) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("mariadb1.username"),
		viper.GetString("mariadb1.password"),
		viper.GetString("mariadb1.addr"),
		viper.GetString("mariadb1.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

/*外部调用的核心方法*/
func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		//Docker: GetDockerDB(), //目前不需要
	}
	log.Infof("初始化完成")
}

func (db *Database) Close() {
	DB.Self.Close()
	//DB.Docker.Close()
}

