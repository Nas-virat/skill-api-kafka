package main

import (
	"context"
	"errors"
	"fmt"
	"gokafka/config"
	"gokafka/database"
	"gokafka/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	db := database.ConnectDB()
	defer db.Close()

	producerConfig, err := config.ProducerKafka()
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producerConfig.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	r := router.NewRouter(db, producerConfig)

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	closeChan := make(chan struct{})

	go func() {
		<-ctx.Done()
		fmt.Println("shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println(err, "err1")
			}
		}

		close(closeChan)
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

	<-closeChan

}
