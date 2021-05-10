package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Name     string
	Age      int
	Birthday int
}
type logger interface {
	Print(v ...interface{})
}

func main() {
	fmt.Println("welcome to aws-rds poc")
	// host := "mydb.cflz2jzzezfj.ap-south-1.rds.amazonaws.com:5432"
	// user := "postgres"
	// password := "postgres"
	// dbname := "postgres"
	// var conn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", host, user, password, dbname)
	// createPgDb()
	db, err := gorm.Open("postgres", "host=my-pg-db.clulvepitg7w.us-east-2.rds.amazonaws.com port=5432 user=postgres dbname=test_db password='postgres' sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("connection established with rds instance", db)

	db.SetLogger(&GormLogger{})
	// db.SetLogger(Sample{})
	db.LogMode(true)
	Formatter := new(log.TextFormatter)
	log.SetFormatter(Formatter)
	Formatter.FullTimestamp = true
	db.AutoMigrate(User{})
	// ===========create operation =============

	user := User{Name: "prashant", Age: 18, Birthday: 9}
	result := db.Create(&user).Debug()
	fmt.Println("error is", result.Error)
	// ===========fetch operation ==========
	var u User
	db.First(&u)
	fmt.Println("data fetched success", u)

}

// GormLogger struct
type GormLogger struct{}

// Print - Log Formatter
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		log.WithFields(
			log.Fields{
				"module":        "gorm",
				"type":          "sql",
				"rows_returned": v[5],
				"src":           v[1],
				"values":        v[4],
				"duration":      v[2],
			},
		).Info(v[3])
	case "log":
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
