package database

import (
	"database/sql"
	"fmt"
	"hvi_server/security"
	"log"
	"time"
	_ "github.com/lib/pq"
)

const CONNECTION_STRING string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

var g_instances[MAX_USERTYPE] *sql.DB

func Init() {
        initConfig()
        
        for i := range MAX_USERTYPE {
                str_connexion := fmt.Sprintf(CONNECTION_STRING,
                                g_configs[i].Host, g_configs[i].Port, g_configs[i].Login, 
                                security.GetSecureStringContent(&g_configs[i].Password), g_configs[i].DatabaseName);
                security.DestroySecureString(&g_configs[i].Password)
                var err error;
                g_instances[i], err = sql.Open("postgres", str_connexion);
                if err != nil {
                        log.Fatalf("Failed to open DB connection: %v", err)
                }
                if err := g_instances[i].Ping(); err != nil {
                        log.Fatalf("Failed to ping DB: %v", err)
                }
                g_instances[i].SetConnMaxLifetime(time.Minute * 5)
                g_instances[i].SetConnMaxIdleTime(time.Minute * 5)
                g_instances[i].SetMaxIdleConns(20)
                g_instances[i].SetMaxOpenConns(10)
        }
	log.Println("Successfully open all database connections")
}

func Close() {
        for i := range MAX_USERTYPE {
                err := g_instances[i].Close()
                if err != nil {
                        log.Fatalf("Error closing database : %v", err)
                }
	}
	log.Println("Successfully close all database connections")
}
