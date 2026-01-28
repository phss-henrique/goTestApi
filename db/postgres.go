package db
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func Connect() *sql.DB{
	dsn := "host=localhost port=5432 user=pedro password=root dbname=pedro sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if error := db.Ping(); error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connection Stablished!")
	return db
}