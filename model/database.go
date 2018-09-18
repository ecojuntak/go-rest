package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

var (
	db  *gorm.DB
	err error
)

func InitDatabase() (err error) {
	dbDriver := viper.GetString("DB_DRIVER")
	var connectionString string

	if dbDriver == "mysql" {
		connectionString = buildMysqlConnectionString()
	} else if dbDriver == "postgres" {
		connectionString = buildPostgresqlConnectionString()
	} else if dbDriver == "sqlite3" {
		connectionString = buildSqliteConnectionString()
	}

	db, err = openConnection(dbDriver, connectionString)

	return
}

func openConnection(dbDriver, connection string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbDriver, connection)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func buildMysqlConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", viper.GetString("DB_USERNAME"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_NAME"))
	return
}

func buildPostgresqlConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", viper.GetString("DB_HOST"), viper.GetString("DB_PORT"), viper.GetString("DB_USERNAME"), viper.GetString("DB_NAME"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_POSTGRES_SSL_MODE"))
	return
}

func buildSqliteConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf(viper.GetString("DB_NAME"))
	return
}
