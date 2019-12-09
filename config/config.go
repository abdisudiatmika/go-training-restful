package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	//DB Database
	DB *gorm.DB
)

func init() {
	InitDB()

}

//InitDB fungsi inisialisasi DB
func InitDB() {
	db, err := gorm.Open("sqlite3", "./db/crud_go.db")
	if err != nil {
		panic(err)
	}
	DB = db
}

//Config untuk menyimpan user
// type Config struct {
// 	DBUsername string
// 	DBPassword string
// 	DBPort     string
// 	DBHost     string
// 	DBName     string
// }

// func initDB() {
// 	config := Config{
// 		DBUsername: "root",
// 		DBPassword: "",
// 		DBPort:     "3306",
// 		DBHost:     "localhost",
// 		DBName:     "dbmvc",
// 	}

// 	connectionString :=
// 		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 			config.DBUsername,
// 			config.DBPassword,
// 			config.DBHost,
// 			config.DBPort,
// 			config.DBName,
// 		)

// 	var err error
// 	db, err := gorm.Open("mysql", connectionString)
// 	if err != nil {
// 		panic(err)
// 	}
// 	DB = db
// }
