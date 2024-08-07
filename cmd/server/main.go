package main

import (
	"go-learn/internal/config"
	"go-learn/internal/database"
	"go-learn/internal/user"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	http.HandleFunc("/users", userHandler.HandleUsers)
	http.HandleFunc("/users/", userHandler.HandleUser)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Responsible for starting the server
// Has DB connection and HTTP Server
