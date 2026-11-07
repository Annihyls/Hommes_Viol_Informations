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

const ConnectionString string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

var gInstances [MaxUserType]Instance

type Instance struct {
	config       *ConfigDBConnection
	dbConnection *sql.DB
}

func GetDatabaseConnection(userType UserType) *sql.DB {
	return gInstances[userType].dbConnection
}

func createConnexion(config *ConfigDBConnection) (*sql.DB, error) {
	strConnection := fmt.Sprintf(ConnectionString,
		config.Host, config.Port, security.GetSecureStringContent(&config.Login),
		security.GetSecureStringContent(&config.Password), config.DatabaseName)
	security.DestroySecureString(&config.Login)
	security.DestroySecureString(&config.Password)
	var dbConnection *sql.DB
	var err error
	dbConnection, err = sql.Open("postgres", strConnection)
	if err != nil {
		return nil, fmt.Errorf("fail to open database connection: %w", err)
	}
	if err := dbConnection.Ping(); err != nil {
		return nil, fmt.Errorf("fail to ping database: %w", err)
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
		return fmt.Errorf("fail to establish connection to database: %w", err)
	}
	return nil
}

func initConnections() error {
	for i := range MaxUserType {
		err := initConnection(&gInstances[i])
		if err != nil {
			for j := i - 1; i >= 0; j-- {
				errfatal := closeConnection(&gInstances[j])
				if errfatal != nil {
					logger.Stderr.Fatalf(
						"Fail to close connection for %s: %v",
						userTypeName[UserType(j)], errfatal)

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
		return fmt.Errorf("error closing database : %w", err)
	}
	return nil
}

func closeConnections() error {
	var errs []error
	for i := range MaxUserType {
		err := closeConnection(&gInstances[i])
		if err != nil {
			errs = append(errs, fmt.Errorf("%s : %w", userTypeName[UserType(i)], err))
		}
	}
	return errors.Join(errs...)
}
