// main.go (Backend)
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	_ "github.com/rs/cors"
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
	initializeMessages()
	fmt.Println("Connected to the database!")
}

func initializeMessages() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL
	);`
	_, erra := db.Exec(createTableQuery)
	if erra != nil {
		return fmt.Errorf("failed to create messages table: %v", erra)
	}

	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		messages := []string{
			"Hello, dude!",
			"Welcome to our new API!",
			"This is me testing out deployments and stuff.",
			"Go is great for building APIs.",
			"React is a fantastic frontend library.",
			"Docker simplifies containerization.",
			"Kubernetes is awesome for orchestration.",
			"PostgreSQL is a powerful database.",
			"Let's build something amazing!",
			"Keep calm and code on.",
			"Error handling is important.",
			"Always test your code.",
			"Documentation is key to great software.",
			"Version control saves lives.",
			"Code reviews help improve quality.",
			"Clean code is easy to read.",
			"Comments should explain why, not what.",
			"Refactor early and often.",
			"Stay curious and keep learning.",
			"Community contributes to growth.",
			"Collaboration leads to success.",
		}

		for _, msg := range messages {
			_, err := db.Exec("INSERT INTO messages (content) VALUES ($1)", msg)
			if err != nil {
				return err
			}
		}
		fmt.Println("Initialized messages")
	}

	return nil
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
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", getMessageHandler).Methods("GET")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsOptions.Handler(router)

	fmt.Println("Backend running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", handler))
}
