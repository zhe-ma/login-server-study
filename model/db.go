package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func (db *Database) Open() {
	DB = &Database{
		Self:   OpenSelfDb(),
		Docker: nil,
	}
}

func (db *Database) Close() {
	if DB.Self != nil {
		DB.Self.Close()
	}

	if DB.Docker != nil {
		DB.Docker.Close()
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

func OpenMysqlDB(username, password, addr, dbName string) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		dbName,
		true,
		"Local")

	log.Debugf("DB connection string: %s", connectionString)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Errorf(err, "Failed to connection DB")
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
}
