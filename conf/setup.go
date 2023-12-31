package conf

import (
	"fmt"
	"log"

	"rutasMap/v2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "lur"
const DB_PASSWORD = "wewDe2323_412323"
const DB_NAME = "busdatabase"
const DB_HOST = "database-1.ctubvgkofn0e.us-east-1.rds.amazonaws.com"

//const DB_USERNAME = "root"
//const DB_PASSWORD = "root"
//const DB_NAME = "testp"
//const DB_HOST = "localhost"

const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() (*gorm.DB, error) {

	return connectDB()
}

func connectDB() (*gorm.DB, error) {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		log.Fatal(err.Error())
	}

	return db, err
}

func Migrate() {
	db, err := connectDB()
	db.AutoMigrate(&models.Point{})
	db.AutoMigrate(&models.Route{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Location{})
	db.AutoMigrate(&models.RouteShow{})
	if err != nil {
		fmt.Println("Error Migrate ")
		log.Fatal(err.Error())
	}

}
