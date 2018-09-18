package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	db  *gorm.DB
	err error
)

func InitDatabase() {
	dbDriver := viper.GetString("DB_DRIVER")
	var connectionString string

	if dbDriver == "mysql" {
		connectionString = buildMysqlConnectionString()
	} else if dbDriver == "postgres" {
		buildPostgresqlConnectionString()
	} else if dbDriver == "sqlite" {
		buildSqliteConnectionString()
	} else if dbDriver == "sqlserver" {
		buildSqlserverConnectionString()
	}

	db, err = openConnection(dbDriver, connectionString)

	if err != nil {
		log.Println(err)
	}
}

func openConnection(dbDriver, connection string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbDriver, connection)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func buildMysqlConnectionString() (connectionString string) {
	connectionString = viper.GetString("DB_USERNAME") + ":" + viper.GetString("DB_PASSWORD") + "@/" + viper.GetString("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	return
}

func buildPostgresqlConnectionString() {
	fmt.Println("building string for postgres")
}

func buildSqliteConnectionString() {
	fmt.Println("building string for sqlite")
}

func buildSqlserverConnectionString() {
	fmt.Println("building string for sqlserver")
}
