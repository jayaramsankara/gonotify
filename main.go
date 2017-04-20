package main

import (
	"github.com/jayaramsankara/gotell"
	"gopkg.in/redis.v5"
	"log"
	"os"
	"strconv"
)

func main() {

	httpHost := "0.0.0.0"

	httpPort := os.Getenv("PORT")

	if httpPort == "" {
		log.Println("$PORT must be set")
		panic("$PORT is not set")
	}
	portNumber, err := strconv.Atoi(httpPort)
	if err != nil {

		log.Println("Invlid port number.", err)
		panic("Invalid port number.")

	}
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Println("$REDIS_URL must be set")
		panic("$REDIS_URL is not set")
	}
	
	redisOptions, rerr := redis.ParseURL(redisURL)
	if rerr != nil {

		log.Println("Invlid redis URL.", rerr)
		panic("Invalid redis URL.")

	}
	err = gotell.InitServer(httpHost, portNumber, redisOptions)

	if err != nil {

		log.Println("Failed to initiate notification service.", err)

	}
}
