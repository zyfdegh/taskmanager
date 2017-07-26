package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func initDB() {
	dbAddr := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "password"
	dbName := "cdn"
	dbParams := "charset=utf8&parseTime=True&loc=Local"

	dbConnString := fmt.Sprintf("%s:%s@%s:%s/%s", dbUser, dbPass, dbAddr, dbPort, dbName)
	if len(dbParams) > 0 {
		dbConnString += "?" + dbParams
	}

	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("open db error: %v\n", err)
	}
	defer db.Close()
}
