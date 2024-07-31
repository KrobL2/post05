package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restfull-server/internal/config"
	"restfull-server/package/handler"
	"time"
)

func main() {
	// Использовать константу из env
	config, err := config.LoadConfig("/restfull-server/config/main.yaml")
	fmt.Print(config == nil)

	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	PORT := ":" + config.HTTPServer.Port

	arguments := os.Args
	if len(arguments) >= 2 {
		PORT = ":" + arguments[1]
	}

	handlers := new(handler.Handler)
	rMux := handlers.InitRoutes()

	s := http.Server{
		Addr:         PORT,
		Handler:      rMux,
		ErrorLog:     nil,
		ReadTimeout:  time.Duration(config.HTTPServer.Timeout),
		WriteTimeout: time.Duration(config.HTTPServer.WriteTimeout),
		IdleTimeout:  time.Duration(config.HTTPServer.IdleTimeout),
	}

	go func() {
		log.Println("Listening to", PORT)

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			return
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	sig := <-sigs
	log.Println("Quitting after signal:", sig)
	time.Sleep(5 * time.Second)
	// s.Shutdown(nil) TODO check if it works
	s.Shutdown(context.Background())
}
