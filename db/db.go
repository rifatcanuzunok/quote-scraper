// db/db.go
package db

import (
	"database/sql"
	"fmt"
	"log"
	"quote-scraper/config"
	"quote-scraper/scraper"

	_ "github.com/lib/pq"
)

// InsertQuotes inserts multiple quotes into the database
func InsertQuotes(db *sql.DB, quotes []scraper.Quote, category string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO quotes (text, author, category) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	for _, quote := range quotes {
		_, err := stmt.Exec(quote.Text, quote.Author, category)
		if err != nil {
			log.Println("Error inserting quote:", err)
			return err
		}
		fmt.Println("Quote inserted successfully:", quote.Text)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// OpenDB opens a connection to the PostgreSQL database
func OpenDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBName))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
