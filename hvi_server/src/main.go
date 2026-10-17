package main

import (
	"context"
	"hvi_server/database"
	"hvi_server/routes"
	"hvi_server/utils/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
        logger.Stdout.Println("Server starting boot...")
        err := database.Init()
        if err != nil {
                logger.Stderr.Printf("%v", err)
        }
        logger.Stdout.Println("Successfully open all database connections")
        routes.RoutesInit()
        logger.Stdout.Println("Routes successfully created")
        srv := &http.Server{
                Addr: ":8080",
        }
        go func() {
                if err := srv.ListenAndServe(); err != nil &&
                        err != http.ErrServerClosed {
                        logger.Stderr.Println("Erreur serveur :", err)
                }
        }()
        // Ton code exécuté pendant que le serveur tourne
        logger.Stdout.Println("Go Server started at: http://localhost:8080")
        
        stop := make(chan os.Signal, 1)
        signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
        <-stop

        // Arrêt propre du serveur
        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()

        if err := srv.Shutdown(ctx); err != nil {
                logger.Stderr.Println("Erreur lors de l'arrêt :", err)
        }

        // Ce code s'exécute après l'arrêt du serveur
        logger.Stdout.Println("Le serveur est arrêté")
        err = database.Close()
        if err != nil {
                logger.Stderr.Printf("%v", err)
        }
	logger.Stdout.Println("Successfully close all database connections")
}
