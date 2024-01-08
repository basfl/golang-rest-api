package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//for datbase connection
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// iniDB()
func init() {
	godotenv.Load()
	username := os.Getenv("DBUSERNAME")
	password := os.Getenv("DBPASSWORD")
	host := os.Getenv("DBHOST")
	schema := os.Getenv("DBSCHEMA")
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	fmt.Println(dataSourceName)
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("***************************")
	fmt.Println("host:", host)
	fmt.Println("schema:", schema)
	log.Println("database successfully configured!")
	fmt.Println("database successfully configured")
	fmt.Println("*********************************")
	createTable()

}

func createTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(767) NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("COULD NOT USERS TABLE")

	}

	createEventTablle := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	`

	_, err = DB.Exec(createEventTablle)
	if err != nil {
		panic("COULD NOT CREATE EVENT TABLE")
	}

	createRegisterationsTable := `

	CREATE TABLE IF NOT EXISTS registerations (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createRegisterationsTable)
	if err != nil {
		panic("COULD NOT CREATE Registeration TABLE")
	}
}
