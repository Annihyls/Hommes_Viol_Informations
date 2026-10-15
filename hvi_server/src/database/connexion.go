package db

import (
	"database/sql"
	"fmt"
	"hvi_server/constants"
	"log"
	"os"
	"strings"
	"time"
	_ "github.com/lib/pq"
)

type UserType int
const (
        POSTGRES UserType = iota
        INTERNAL
        EXTERNAL
)

type DBConnectionConfig struct {
	UserType     UserType
	DatabaseName string
	Host         string
	Port         string
	Login        string
	Password     string
}

var g_instances[3] *sql.DB
var g_configs[3]   DBConnectionConfig

func getLoginPasswordFromUserType(userType UserType) (string, string) {
        var login    string
        var password string
        switch userType {
                case POSTGRES:
                login = os.Getenv(cstenv.USER_POSTGRES_LOGIN)
                password = os.Getenv(cstenv.USER_POSTGRES_PASSWORD)
                case INTERNAL:
                login = os.Getenv(cstenv.USER_INTERNAL_LOGIN)
                password = os.Getenv(cstenv.USER_INTERNAL_PASSWORD)
                default:
                login = os.Getenv(cstenv.USER_EXTERNAL_LOGIN)
                password = os.Getenv(cstenv.USER_EXTERNAL_PASSWORD)
        }
        return login, password
}

func createConnexionConfigWithUserType(userType UserType) DBConnectionConfig {
        dbCoCfg := DBConnectionConfig {
                UserType: userType,
                Host: os.Getenv(cstenv.DB_HOST),
                Port: os.Getenv(cstenv.DB_PORT),
                DatabaseName: os.Getenv(cstenv.DB_NAME),
        }
        dbCoCfg.Login, dbCoCfg.Password = getLoginPasswordFromUserType(dbCoCfg.UserType)
        checkIfConfigIsConform(&dbCoCfg)
        return dbCoCfg
}

func isStringEmpty(str string) bool {
        return len(strings.TrimSpace(str)) <= 0
}

func checkIfConfigIsConform(cfg *DBConnectionConfig) {
        const msg_error string = "is not defined in your environment. Have you setup your .env file ?"
        if isStringEmpty(cfg.Host) {
                log.Fatalf("%s %s", cstenv.DB_HOST, msg_error)
        } else if isStringEmpty(cfg.Port) {
                log.Fatalf("%s %s", cstenv.DB_PORT, msg_error)
        } else if isStringEmpty(cfg.DatabaseName) {
                log.Fatalf("%s %s", cstenv.DB_NAME, msg_error)
        } else if isStringEmpty(cfg.Login) {
                log.Fatalf("%s %s", cstenv.USER_INTERNAL_LOGIN, msg_error)
        } else if isStringEmpty(cfg.Password) {
                log.Fatalf("%s %s", cstenv.USER_INTERNAL_PASSWORD, msg_error)
        }
}

func Init() {
        g_configs[EXTERNAL] = createConnexionConfigWithUserType(EXTERNAL)
        g_configs[INTERNAL] = createConnexionConfigWithUserType(INTERNAL)
        g_configs[POSTGRES] = createConnexionConfigWithUserType(POSTGRES)

        for i := range g_configs {
                dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
                                g_configs[i].Host, g_configs[i].Port, g_configs[i].Login, 
                                g_configs[i].Password, g_configs[i].DatabaseName);
                var err error;
                g_instances[i], err = sql.Open("postgres", dsn);
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
        for i := range g_instances {
                err := g_instances[i].Close()
                if err != nil {
                        log.Fatalf("Error closing database : %v", err)
                }
	}
	log.Println("Successfully close all database connections")
}
