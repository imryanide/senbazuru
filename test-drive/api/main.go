// main.go (Backend)
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var db *sql.DB

func connectDB() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	fmt.Println("Connected to the database!")
}

func getMessageHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, content FROM messages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, message)
	}
	json.NewEncoder(w).Encode(messages)
}

func main() {
	connectDB()

	http.HandleFunc("/api/messages", getMessageHandler)
	fmt.Println("Backend running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
