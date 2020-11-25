package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Connection struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
	dbConnStr  string
}

var conn Connection

func TestConn() *gorm.DB {
	setTestCredentials()

	db, err := gorm.Open(postgres.Open(conn.dbConnStr), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func setTestCredentials() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	conn.dbHost = os.Getenv("POSTGRES_TEST_HOST")
	conn.dbPort = os.Getenv("POSTGRES_TEST_PORT")
	conn.dbUser = os.Getenv("POSTGRES_TEST_USER")
	conn.dbPassword = os.Getenv("POSTGRES_TEST_PASSWORD")
	conn.dbName = os.Getenv("POSTGRES_TEST_DB")

	conn.dbConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conn.dbHost, conn.dbPort, conn.dbUser, conn.dbPassword, conn.dbName)
}

func Conn() *gorm.DB {
	setCredentials()
	db, err := gorm.Open(postgres.Open(conn.dbConnStr), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func setCredentials() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	conn.dbHost = os.Getenv("POSTGRES_HOST")
	conn.dbPort = os.Getenv("POSTGRES_PORT")
	conn.dbUser = os.Getenv("POSTGRES_USER")
	conn.dbPassword = os.Getenv("POSTGRES_PASSWORD")
	conn.dbName = os.Getenv("POSTGRES_DB")

	conn.dbConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conn.dbHost, conn.dbPort, conn.dbUser, conn.dbPassword, conn.dbName)
}
