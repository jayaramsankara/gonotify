package main

import (
	"github.com/jayaramsankara/gotell"
	"gopkg.in/redis.v5"
	"log"
	"os"
	"strconv"
	"github.com/fsnotify/fsnotify"
	"github.com/kardianos/osext"
	"fmt"
	"syscall"
	
)

// watchSelf watches for changes in the main binary and hot-swaps itself for the newly
// built binary file
func watchSelf() (chan struct{}, error) {

    // Retrieve file info for the currently running program
    file, err := osext.Executable()
    if err != nil {
        return nil, err
    }

    // Initialize and prepare a new file watcher
    fmt.Printf("watching %q\n", file)
    w, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err
    }

    done := make(chan struct{})
    go func() {
        for {
            select {
                case e := <-w.Events:
                    // Events mean changes
                    fmt.Printf("watcher received: %+v", e)

                    // Replace the running system call with a new call
                    // to our newly combined binary
                    err := syscall.Exec(file, os.Args, os.Environ())
                    if err != nil {
                        fmt.Errorf("%#v", err)
                    }

                case err := <-w.Errors:
                    // Print out errors as they occur
                    fmt.Printf("watcher error: %+v", err)

                case <-done:
                    // If we ever close the watcher, log it 
                    fmt.Printf("watcher shutting down")
                    return
            }
        }
    }()

    // Add our running file to be watched
    err = w.Add(file)
    if err != nil {
        return nil, err
    }
    return done, nil
}

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
	log.Println("Starting notification service. Watching the exec now..", err)
	watchSelf()
	
	err = gotell.InitServer(httpHost, portNumber, redisOptions)

	if err != nil {

		log.Println("Failed to initiate notification service.", err)

	} 
}
