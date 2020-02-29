package Data

import (
	"learngolang/Constants"
	"learngolang/Models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateUser(param Models.User) {
	db, err := gorm.Open("postgres", " host="+Constants.Host+" port="+Constants.Port+" user="+Constants.User+" dbname="+Constants.Dbname+" password="+Constants.Password)

	if err != nil {
		panic(err.Error())

	} else {
		db.Create(&param)
	}
	defer db.Close()
}

func GetUserByEmailOrPhoneNumber(param Models.User) Models.User {

	var checkuser Models.User
	db, err := gorm.Open("postgres", " host="+Constants.Host+" port="+Constants.Port+" user="+Constants.User+" dbname="+Constants.Dbname+" password="+Constants.Password)
	db.Where("Email = ? Or Phone = ?", param.Email, param.Phone).First(&checkuser)
	if err != nil {
		panic(err.Error())

	} else {

	}
	defer db.Close()
	return checkuser

}
