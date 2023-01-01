package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rxxcc/nigeria-uni/router"
)

func main() {
	// load the env file
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatalf("error loading env file %s", err)
	}

	// getting our env
	port := os.Getenv("PORT")

	log.Println("server running at port", port)

	route := router.NEW()

	if err := http.ListenAndServe(":"+port, route); err != nil {
		log.Fatal(err)
	}
}
