package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Conn       *gorm.DB
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
	dbConnStr  string
}

func (pg *Postgres) SetPostgres() {
	pg.setCredentials()

	pg.setConn()

}

func (pg *Postgres) setCredentials() {
	pg.dbName = os.Getenv("POSTGRES_DB")
	pg.dbHost = os.Getenv("POSTGRES_HOST")
	pg.dbPort = os.Getenv("POSTGRES_PORT")
	pg.dbUser = os.Getenv("POSTGRES_USER")
	pg.dbPassword = os.Getenv("POSTGRES_PASSWORD")
	pg.dbConnStr = fmt.Sprintf(`host=%s 
								port=%s 
								user=%s 
								password=%s 
								dbname=%s 
								sslmode=disable`,
		pg.dbHost, pg.dbPort, pg.dbUser, pg.dbPassword, pg.dbName)
}

func (pg *Postgres) setConn() {
	conn, err := gorm.Open(postgres.Open(pg.dbConnStr), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	pg.Conn = conn
}
