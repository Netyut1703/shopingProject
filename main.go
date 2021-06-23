package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/netyut1703/shopingProject/routes"
	"github.com/netyut1703/shopingProject/server"
	"github.com/netyut1703/shopingProject/store"
)

func main() {

	// read env varianles from file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// get access to store
	fmt.Println("main.go: Creating the store...")
	store := store.New()

	// setup message queue
	// fmt.Println("main.go: Setting up message queue (kafka)...")
	//queue := queue.New()

	// connect to s3
	s3Store := s3.New()

	// create server instance by passing the store
	fmt.Println("main.go: Creating the server...")
	server := server.New(store, nil, s3Store)

	// init routes and start the server
	fmt.Println("main.go: Initializing the routes...")
	routes := routes.InitRoutes(server)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Content-Disposition"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE"})

	// gracefull server implementation pending
	fmt.Printf("main.go:Http server up on PORT: %s\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handlers.CORS(originsOk, headersOk, methodsOk)(routes)))

}
