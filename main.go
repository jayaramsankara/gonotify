package main

import (
	"os"
	"log"
	"strconv"
	"gopkg.in/redis.v3"
	"github.com/jayaramsankara/gotell"
)

func main() {
	
	httpHost := "0.0.0.0"
	
	httpPort := os.Getenv("PORT")
	
    if httpPort == "" {
        log.Println("$PORT must be set")
		panic("$PORT is not set")
    }
	portNumber , err := strconv.Atoi(httpPort)
	if err != nil {

		log.Println("Invlid port number.", err)
		panic("Invalid port number.")

	}
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
        log.Println("$REDIS_URL must be set")
		panic("$REDIS_URL is not set")
    }
	
	redisOptions := &redis.Options{
				Addr:        os.Getenv("REDIS_URL"),
				Password:    "", // no password set
				DB:          0,  // use default DB
				PoolTimeout: 3,  // Pool timeout
				MaxRetries:  3,
				PoolSize:    50,
			}
	err = gotell.InitServer(httpHost, portNumber, redisOptions)

	if err != nil {

		log.Println("Failed to initiate notification service.", err)

	}
}
