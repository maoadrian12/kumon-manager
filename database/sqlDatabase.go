package database

import (
	"context"
	"fmt"

	"log"
	"os"

	models "kumondatabase.com/models"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/jackc/pgx/v4/pgxpool"

	//"github.com/jackc/pgx/v5"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var Connection *pgxpool.Conn
var Db *sql.DB

func Testing() {
	// Capture connection properties.
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "postgres", Password, "kumon_data", "9292")
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	Database.AutoMigrate(&models.Parents{}, &models.Student{})
	database_url := fmt.Sprintf("postgres://postgres:%s@localhost:9292/kumon_data?sslmode=disable", Password)
	var err1 error
	Db, err1 = sql.Open("postgres", database_url)
	if err1 != nil {
		log.Fatal(err1)
	}

	pool, err := pgxpool.Connect(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Connection, err = pool.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Error getting connection from pool: %s\n", err)
	}
}
