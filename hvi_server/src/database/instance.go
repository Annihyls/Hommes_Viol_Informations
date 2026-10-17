package database

import (
	"database/sql"
	"fmt"
	"hvi_server/security"
	"hvi_server/utils/logger"
	"time"

	_ "github.com/lib/pq"
)

const CONNECTION_STRING string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

var g_instances [MAX_USERTYPE]Instance 

type Instance struct {
        config *ConfigDBConnection
        dbConnection *sql.DB
}

func GetDatabaseConnection(userType UserType) *sql.DB {
	return g_instances[userType].dbConnection;
}

func createConnexion(config *ConfigDBConnection) (*sql.DB, error) {
        str_connexion := fmt.Sprintf(CONNECTION_STRING,
                        config.Host, config.Port, config.Login, 
                        security.GetSecureStringContent(&config.Password), config.DatabaseName);
        security.DestroySecureString(&config.Password)
        var dbConnection *sql.DB
        var err error;
        dbConnection, err = sql.Open("postgres", str_connexion);
        if err != nil {
                return nil, fmt.Errorf("Fail to open DB connection: %w", err)
        }
        if err := dbConnection.Ping(); err != nil {
                return nil, fmt.Errorf("Fail to ping DB: %w", err)
        }
        dbConnection.SetConnMaxLifetime(time.Minute * 5)
        dbConnection.SetConnMaxIdleTime(time.Minute * 5)
        dbConnection.SetMaxIdleConns(20)
        dbConnection.SetMaxOpenConns(10)
        return dbConnection, nil
}

func initConnection(instance *Instance) error {
        var err error
        instance.dbConnection, err = createConnexion(instance.config)
        if err != nil {
                return fmt.Errorf("Fail to establish connection to database: %w", err)
        }
        return nil
}

func initConnections() error {
        err := initConnection(&g_instances[POSTGRES])
        if err != nil {
                return fmt.Errorf("Postgres user: %w", err)
        }
        err = initConnection(&g_instances[INTERNAL])
        if err != nil {
                errfatal := closeConnection(&g_instances[POSTGRES])
                if errfatal != nil {
                        logger.Stderr.Fatalf("Fail to close connection Postgres: %v", errfatal)
                }
                return fmt.Errorf("Internal user: %w", err)
        }
        err = initConnection(&g_instances[EXTERNAL])
        if err != nil {
                errfatal := closeConnection(&g_instances[POSTGRES])
                if errfatal != nil {
                        logger.Stderr.Fatalf("Fail to close connection Postgres: %v", errfatal)
                }
                errfatal = closeConnection(&g_instances[INTERNAL])
                if errfatal != nil {
                        logger.Stderr.Fatalf("Fail to close connection Internal: %v", errfatal)
                }
                return fmt.Errorf("External user: %w", err)
        }
        return nil
}

func closeConnection(instance *Instance) error {
        err := instance.dbConnection.Close()
        if err != nil {
                return fmt.Errorf("Error closing database : %w", err)
        }
        return nil
}

func closeConnections() error {
        err := closeConnection(&g_instances[POSTGRES])
        if err != nil {
                return fmt.Errorf("Postgres : %w", err)
        }
        err = closeConnection(&g_instances[INTERNAL])
        if err != nil {
                return fmt.Errorf("Internal : %w", err)
        }
        err = closeConnection(&g_instances[EXTERNAL])
        if err != nil {
                return fmt.Errorf("External : %w", err)
        }
        return nil
}
