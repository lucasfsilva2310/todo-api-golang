package main

import (
	"database/sql"
	"go-learn/internal/config"
	"go-learn/internal/database"
	"go-learn/internal/user"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	var db *sql.DB = dbConnect(cfg)
	defer db.Close()

	log.Println("Connected to database.")

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	http.HandleFunc("/users", userHandler.HandleUsers)
	http.HandleFunc("/users/", userHandler.HandleUser)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dbConnect(cfg *config.Config) *sql.DB {
	for i := 1; i <= 10; i++ {
		log.Println("Connecting to database, attempt:", i)
		connectedDB, err := database.Connect(cfg)

		if err != nil {
			log.Printf("Failed to connect to database: %v", err)
		}

		err = connectedDB.Ping()
		if err == nil {
			return connectedDB
		}
		log.Println("Waiting for database to be ready...")
		time.Sleep(5 * time.Second)

	}
	log.Println("Could not connect to database!")
	return nil
}
