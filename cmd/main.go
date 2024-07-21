package main

import (
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
	config, err := config.LoadConfig("./config/main.yaml") // Вынести в константу путь до файла конфигурации
	// fmt.Print(config)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	PORT := ":" + config.Server.Port

	arguments := os.Args
	fmt.Println(arguments)
	if len(arguments) >= 2 {
		PORT = ":" + arguments[1]
	}

	handlers := new(handler.Handler)
	rMux := handlers.InitRoutes()

	s := http.Server{
		Addr:         PORT,
		Handler:      rMux,
		ErrorLog:     nil,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout),
		WriteTimeout: time.Duration(config.Server.WriteTimeout),
		IdleTimeout:  time.Duration(config.Server.IdleTimeout),
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
	s.Shutdown(nil)
}
