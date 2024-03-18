package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type parents struct {
	Username string `gorm:"primaryKey"`
	Name     string
	Pass     string
}

func main() {
	// Capture connection properties.
	var dsn = "host=localhost user=postgres password=4B-R05miyo dbname=kumon_database port=9292 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database_url := "postgres://postgres:4B-R05miyo@localhost:9292/kumon_database"
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	var parent parents
	var user = db.Take(&parent, "username", "maoadrian12")
	/*var username string
	var name string
	var pass string
	err = conn.QueryRow(context.Background(), "select username, name, pass from parents").Scan(&username, &name, &pass)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}*/
	//fmt.Println(username, name, pass)

	fmt.Printf("username: %s\n", parent.Username)
	fmt.Printf("%d rows effected\n", user.RowsAffected)
}
