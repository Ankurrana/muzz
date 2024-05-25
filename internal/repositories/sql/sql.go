package sql

import (
	"flag"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string  `json:"name"`
	Email        string  `json:"email" gorm:"unique;not null"`
	Age          int     `json:"age"`
	Gender       string  `json:"gender"`
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
	PasswordHash string  `json:"passwordhash"`
}

type Match struct {
	gorm.Model
	UserA int `gorm:"index"`
	UserB int `gorm:"index"`
}

type Swipe struct {
	gorm.Model
	FromUser int `gorm:"uniqueIndex:fromto;not null"`
	ToUser   int `gorm:"uniqueIndex:fromto;not null"`
	Swipe    bool
}

type UserSummary struct {
	UserId     int `gorm:"index;unique;primarykey"`
	SwipeCount int
	MatchCount int
}

func Connection() (*gorm.DB, error) {
	var dbhost, dbusername, dbname, dbpassword, port string
	flag.StringVar(&dbusername, "dbusername", "postgres", "database username")
	flag.StringVar(&dbhost, "dbhost", "localhost", "database host")
	flag.StringVar(&dbname, "dbname", "postgres", "database name")
	flag.StringVar(&dbpassword, "dbpassword", "postgres", "database password")
	flag.StringVar(&port, "dbport", "5432", "5432")
	flag.Parse()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbhost, dbusername, dbpassword, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error:unable to connect to the db %s", err)
	}
	err2 := db.AutoMigrate(&User{}, &Match{}, &Swipe{}, &UserSummary{})
	if err2 != nil {
		fmt.Print(err2)
	}
	return db, err
}
