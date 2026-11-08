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

type Instance struct {
	config       ConfigDBConnection
	dbConnection *sql.DB
}

func createConnexion(config ConfigDBConnection) (*sql.DB, error) {
	strConnection := fmt.Sprintf(ConnectionString,
		config.Host, config.Port, config.Login,
		security.GetSecureStringContent(&config.Password), config.DatabaseName)
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

func newInstance(dbConfig ConfigDBConnection) (Instance, error) {
	var err error
	var instance Instance
	instance.config = dbConfig
	instance.dbConnection, err = createConnexion(instance.config)
	if err != nil {
		return instance, fmt.Errorf("fail to establish connection to database: %w", err)
	}
	return instance, nil
}

func initConnections(dbConfigs []ConfigDBConnection) ([]Instance, error) {
	instances = make([]Instance, 3)
	for i := range MaxUserType {
		var err error
		instances[i], err = newInstance(dbConfigs[i])
		if err != nil {
			for j := i - 1; i >= 0; j-- {
				errfatal := closeConnection(instances[j])
				if errfatal != nil {
					logger.Stderr.Fatalf(
						"Fail to close connection for %s: %v",
						userTypeName[UserType(j)], errfatal)

				}
			}
			return instances, fmt.Errorf("%s: %w", userTypeName[UserType(i)], err)
		}
	}
	return instances, nil
}

func closeConnection(instance Instance) error {
	err := instance.dbConnection.Close()
	if err != nil {
		return fmt.Errorf("error closing database : %w", err)
	}
	return nil
}

func closeConnections(instances []Instance) error {
	var errs []error
	for i := range MaxUserType {
		err := closeConnection(instances[i])
		if err != nil {
			errs = append(errs, fmt.Errorf("%s : %w", userTypeName[UserType(i)], err))
		}
	}
	return errors.Join(errs...)
}
