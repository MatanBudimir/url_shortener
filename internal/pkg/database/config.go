package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// DB represents database connection variable
var (
	DB                                   *gorm.DB
	err                                  error
	user, password, database, host, port string
)

// Configure establishes connections with the database
func Configure() {
	env()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Database connection established")
}

// env loads and reads the environmental variables
func env() {
	err = godotenv.Load("../../.env")

	if err != nil {
		log.Fatal(err)
		return
	}

	user = os.Getenv("DATABASE_USER")
	password = os.Getenv("DATABASE_PASSWORD")
	database = os.Getenv("DATABASE_NAME")
	host = os.Getenv("DATABASE_HOST")
	port = os.Getenv("DATABASE_PORT")

	if user == "" || password == "" || database == "" || host == "" || port == "" {
		log.Fatal("Please configure environmental file.")
	}
}
