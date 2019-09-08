package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Here is just a test print for main")
	fmt.Println("Starting up the router handler")

	router := mux.NewRouter()
	router.HandleFunc("/", testHandler)

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)

	fmt.Println("Shutdown")
	os.Exit(0)
}
