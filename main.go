package main

import (
	"github.com/joho/godotenv"
	"golang-ddd-demo/middleware"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	middleware.InitServer().RunServer()
}
