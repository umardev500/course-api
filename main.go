package main

import (
	"context"
	"course-api/application"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app := application.App{}
	app.Start(context.TODO())
}
