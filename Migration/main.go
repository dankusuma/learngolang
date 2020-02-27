package main

import (
	"go-rest-api/Constants"
	"go-rest-api/Models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func MigrateDataUser() {
	db, err := gorm.Open("postgres", " host="+Constants.Host+" port="+Constants.Port+" user="+Constants.User+" dbname="+Constants.Dbname+" password="+Constants.Password)

	if err != nil {
		panic(err.Error())

	} else {
		db.AutoMigrate(&Models.User{})
	}
	defer db.Close()
}

func main() {
	MigrateDataUser()
}
