package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL() (*sql.DB, error) {
	// Environment variable
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), "3306", os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for {
		err := db.Ping()
		if err == nil {
			err = createTodoTable(db)
			if err != nil {
				fmt.Println("Table already exists...", err)
			}

			err = seedTodoList(db)
			if err != nil {
				fmt.Println("Seeding error,", err)
			}

			break
		}
		fmt.Println(dsn)
		fmt.Println("Connecting to db...")
		time.Sleep(1 * time.Second)
	}

	return db, nil
}

// Simple migration
func createTodoTable(db *sql.DB) error {
	createTodoQuery := `CREATE TABLE IF NOT EXISTS todos(
        ID int AUTO_INCREMENT PRIMARY KEY,
        TASK VARCHAR(50),
        STATUS BOOL
    )`

	_, err := db.Exec(createTodoQuery)
	if err != nil {
		return err
	}

	return nil
}

// Seed tables
func seedTodoList(db *sql.DB) error {
	getRows := `SELECT * FROM todos LIMIT 1`

	rows, err := db.Query(getRows)
	if rows.Next() {
		return fmt.Errorf("Table already seeded")
	}
	if err != nil {
		return err
	}

	insertQuery :=
		`INSERT INTO todos(task, status) 
    VALUES ('Join Monthly Cloud', false),
           ('Join Weekly Cloud', true);`

	_, err = db.Exec(insertQuery)
	if err != nil {
		return err
	}

	return nil
}
