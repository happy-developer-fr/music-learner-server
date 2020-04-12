package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

var DataBase *gorm.DB

func OpenConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	DataBase, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = DataBase.DB().Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	fmt.Println("Successfully connected!")
}
