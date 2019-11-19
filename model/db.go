package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func (db *Database) Open() {
	db = &Database{
		Self:  OpenMysqlDB(),
		Doker: nil,
	}
}

func (db *Database) Close() {
	if db.Self != nil {
		db.Self.Close()
	}

	if db.Docker != nil {
		db.Docker.Close()
	}

	// DB.Docker.Close()
}

func OpenSelfDb() *gorm.DB {
	return OpenMysqlDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func OpenDockerDb() *gorm.DB {
	return OpenMysqlDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func OpenMysqlDB(username, password, addr, db_name string) *gorm.DB {
	connection_string := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		db_name,
		true,
		"Local")

	log.Debugf("DB connection string: %s", connection_string)

	db, err := gorm.Open("mysql", connection_string)
	if err != nil {
		log.Errorf(err, "Failed to connection DB")
	}

	setupDB()

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
}