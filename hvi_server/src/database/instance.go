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

var g_instances [MAX_USERTYPE]Instance 

type Instance struct {
        config ConfigDBConnection
        dbConnection *sql.DB
}

func GetDatabaseConnection(userType UserType) *sql.DB {
	return g_instances[userType].dbConnection;
}

func createConnexion(config ConfigDBConnection) *sql.DB {
        str_connexion := fmt.Sprintf(CONNECTION_STRING,
                        config.Host, config.Port, config.Login, 
                        security.GetSecureStringContent(&config.Password), config.DatabaseName);
        security.DestroySecureString(&config.Password)
        var dbConnection *sql.DB
        var err error;
        dbConnection, err = sql.Open("postgres", str_connexion);
        if err != nil {
                log.Fatalf("Failed to open DB connection: %v", err)
        }
        if err := dbConnection.Ping(); err != nil {
                log.Fatalf("Failed to ping DB: %v", err)
        }
        dbConnection.SetConnMaxLifetime(time.Minute * 5)
        dbConnection.SetConnMaxIdleTime(time.Minute * 5)
        dbConnection.SetMaxIdleConns(20)
        dbConnection.SetMaxOpenConns(10)
        return dbConnection
}

func initConnection(instance *Instance) {
        instance.dbConnection = createConnexion(instance.config)
}

func initConnections() {
        initConnection(&g_instances[POSTGRES])
        initConnection(&g_instances[INTERNAL])
        initConnection(&g_instances[EXTERNAL])
}

func closeConnection(instance Instance) {
        err := instance.dbConnection.Close()
        if err != nil {
                log.Fatalf("Error closing database : %v", err)
        }
}

func closeConnections() {
        closeConnection(g_instances[POSTGRES])
        closeConnection(g_instances[INTERNAL])
        closeConnection(g_instances[EXTERNAL])
}
