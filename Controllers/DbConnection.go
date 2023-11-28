package Controllers

import (
	"database/sql"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/go-gorp/gorp"
)

func DbConnection() *gorp.DbMap {

	// errEnv := godotenv.Load(".env")
	// if errEnv != nil {
	// 	log.Error("failed to load env file")
	// }

	// dbUser := os.Getenv("DB_USERNAME")
	// dbPass := os.Getenv("DB_PASSWORD")
	// dbDriver := os.Getenv("DB_CONNECTION")
	// dbName := os.Getenv("DB_DATABASE")

	// dbTcp := "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+dbTcp+dbName+"?charset=utf8&parseTime=True")

	db, err := sql.Open("mysql", "root:choice123@tcp(10.2.30.101:3306)/std")
	if err != nil {
		panic(err.Error())
	}

	log.Info("db Error: ", err)
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()

	return dbmap

}

func gormDB() (*gorm.DB, error) {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Error("failed to load env file")
	}
	dbDriver := os.Getenv("DB_CONNECTION")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbTcp := "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"

	db, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+dbTcp+dbName+"?charset=utf8&parseTime=True")

	//	db, err := gorm.Open("mysql", "ob-ur:Choice@098@tcp(20.1.50.194:3306)/uat2OnBoardDb")
	if err != nil {
		return nil, err
	}
	return db, nil
}
