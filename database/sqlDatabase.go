package database

import (
	"context"
	"fmt"
	"os"

	models "kumondatabase.com/models"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var Connection *pgx.Conn

func Testing() {
	// Capture connection properties.
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "postgres", Password, "kumon_database", "9292")
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database_url := fmt.Sprintf("postgres://postgres:%s@localhost:9292/kumon_database", Password)
	Connection, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Database.AutoMigrate(&models.Parents{})
	Connection.Begin(context.Background())
	//defer conn.Close(context.Background())
	//var parent parents
	//var user = Db.Take(&parent, "username", "maoadrian12")
	/*var username string
	var name string
	var pass string
	err = conn.QueryRow(context.Background(), "select username, name, pass from parents").Scan(&username, &name, &pass)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}*/
	//fmt.Println(username, name, pass)

	//fmt.Printf("username: %s\n", parent.Username)
	//fmt.Printf("%d rows effected\n", user.RowsAffected)
}
