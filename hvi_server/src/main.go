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
        "hvi_server/configserver"
)

func main() {
        logger.Stdout.Println("Server starting boot...")
        if err := config.Load(); err != nil {
                logger.Stderr.Printf("%v\n", err)
        }
        logger.Stdout.Println("Successfully load config file")
        if err := database.Init(); err != nil {
                panic(err)
        }
        logger.Stdout.Println("Successfully open all database connections")
        routes.Init()
        logger.Stdout.Println("Routes successfully created")
        srv := &http.Server{
                Addr: "[::1]:8080",
        }
        
        go func() {
                if err := srv.ListenAndServe(); err != nil &&
                        err != http.ErrServerClosed {
                        logger.Stderr.Printf("%v", err)
                        return
                }
        }()
        logger.Stdout.Println("Go Server started at: http://[::1]:8080")
        
        stop := make(chan os.Signal, 1)
        signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
        <-stop

        ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
        defer cancel()

        if err := srv.Shutdown(ctx); err != nil {
                panic(err)
        }

        if err := database.Close(); err != nil {
                panic(err)
        }
	logger.Stdout.Println("Successfully close all database connections")
	logger.Stdout.Println("Server stopped")
}
