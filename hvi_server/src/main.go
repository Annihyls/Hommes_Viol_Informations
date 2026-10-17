package main

import (
	"hvi_server/database"
	"hvi_server/routes"
	"net/http"
	"time"
	"context"
	"os"
	"os/signal"
	"syscall"
	"log"
)

func main() {
        log.Println("Server starting boot...")
        database.Init()
        log.Println("Successfully open all database connections")
        routes.RoutesInit()
        log.Println("Routes successfully created")
        srv := &http.Server{
                Addr: ":8080",
        }
        go func() {
                if err := srv.ListenAndServe(); err != nil &&
                        err != http.ErrServerClosed {
                        log.Println("Erreur serveur :", err)
                }
        }()
        // Ton code exécuté pendant que le serveur tourne
        log.Println("Go Server started at: http://localhost:8080")
        
        stop := make(chan os.Signal, 1)
        signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
        <-stop

        // Arrêt propre du serveur
        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()

        if err := srv.Shutdown(ctx); err != nil {
                log.Println("Erreur lors de l'arrêt :", err)
        }

        // Ce code s'exécute après l'arrêt du serveur
        log.Println("Le serveur est arrêté")
        database.Close()
	log.Println("Successfully close all database connections")
}
