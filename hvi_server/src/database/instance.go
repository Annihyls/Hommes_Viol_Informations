package database

import (
	"database/sql"
	"errors"
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
                return nil, fmt.Errorf("Fail to open database connection: %w", err)
        }
        if err := dbConnection.Ping(); err != nil {
                return nil, fmt.Errorf("Fail to ping database: %w", err)
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
        for i := range MAX_USERTYPE {
                err := initConnection(&g_instances[i])
                if err != nil {
                        for j := i - 1; i >= 0; j-- {
                                errfatal := closeConnection(&g_instances[j])
                                if errfatal != nil {
                                        logger.Stderr.Fatalf("Fail to close connection for %s: %v", userTypeName[UserType(j)], errfatal)
                                        
                                }
                        }
                        return fmt.Errorf("%s: %w", userTypeName[UserType(i)], err)
                }
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
        var errs []error
        for i := range MAX_USERTYPE {
                err := closeConnection(&g_instances[i])
                if err != nil {
                        errs = append(errs, fmt.Errorf("%s : %w",userTypeName[UserType(i)], err))
                }
        }
        return errors.Join(errs...)
}
