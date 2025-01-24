package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("Last connection saved: [%s]", timestamp)

	_, err := db.Exec("INSERT INTO messages (content) VALUES (?)", message)
	if err != nil {
		log.Printf("Error inserting data into database: %v", err)
		http.Error(w, "Error inserting data into the database", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Message inserted successfully: %s\n", message)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error getting hostname: %v", err)
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT id, content FROM messages")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Error querying data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprintf(w, "Backend Host: %s\n", hostname)
	fmt.Fprintf(w, "Data in the database:\n")
	for rows.Next() {
		var id int
		var content string
		if err := rows.Scan(&id, &content); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "ID: %d, Message: %s\n", id, content)
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Missing environment variables for database configuration")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id INT AUTO_INCREMENT PRIMARY KEY,
		content TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatalf("Error creating the table: %v", err)
	}

	http.HandleFunc("/", dataHandler)
	http.HandleFunc("/save", saveHandler)

	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
