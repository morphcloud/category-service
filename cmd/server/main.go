package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/morphcloud/order-service/internal/database"
	"github.com/morphcloud/order-service/internal/routes"
)

func waitForShutdown(srv http.Server, l *log.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Graceful shutdown:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	l := log.New(os.Stdout, "API ", log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		l.Fatalln(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Port is not set.")
	}

	mongoClient, err := database.NewMongoClient(ctx)
	if err != nil {
		l.Fatalln(err)
	}
	defer mongoClient.Disconnect(ctx)

	router := mux.NewRouter()

	routes.MapURLPathsToHandlers(router, mongoClient, l)

	srv := http.Server{
		Addr:              ":"+port,
		Handler:           router,
		ErrorLog:          l,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			l.Fatalln(err)
		}
	}()
	l.Println("Service has been started")

	waitForShutdown(srv, l)
}
