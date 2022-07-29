package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goApi/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadDb() *sql.DB {
	//laod the app's environment
	err := godotenv.Load()

	//flag an err if the env refuses to load
	if err != nil {
		log.Fatalf("Error loading %s", err)
	} else {
		log.Println("env file loaded")
	}

	//load the env variables
	username := os.Getenv("APP_DB_USERNAME")
	password := os.Getenv("APP_DB_PASSWORD")
	host := os.Getenv("APP_DB_HOST")
	port := os.Getenv("APP_DB_PORT")
	dbname := os.Getenv("APP_DB_NAME")

	//create the connection string
	connects := fmt.Sprintf("host=%s port=%s username=%s "+"password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	fmt.Println(connects)

	//connect the database
	db, err := sql.Open("postgres", connects)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Dabase connected successfully...")

	return db
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	books := []models.Book{}

	db := LoadDb()

	sqlStament := `SELECT * FROM books`

	rows, err := db.Query(sqlStament)

	if err != nil {
		log.Println(fmt.Sprintf("error occured while doing this :%s", err))
	}

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Year, &book.CreatedAt)

		if err != nil {
			log.Println(fmt.Sprintf("error occured while doing this :%s", err))
		}

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)

	defer db.Close()
	defer rows.Close()

}
