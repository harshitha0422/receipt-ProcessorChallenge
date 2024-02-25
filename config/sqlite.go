package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

//To find any error in Database
func checkErr(err error, sucess_message string) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sucess_message)
	}
}

//Db connection establishment through GORM,GORM officially supports databases MySQL
func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:prepboost#123@tcp(prepboost-prod.cppktldlc6tz.ap-south-1.rds.amazonaws.com:3306)/prepboost?parseTime=true")
	fmt.Println(err)
	checkErr(err, "Connected to AWS DB successfully")

	// DBMigrate(db)
	return db
}
