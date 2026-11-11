package database

import (
	"database/sql"
	"fmt"
	"hvi_server/security"
	"time"

	_ "github.com/lib/pq"
)

const ConnectionString string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

type DBInstance struct {
	config       DBConfig
	dbConnection *sql.DB
}

func createDBInstance(dbConfig DBConfig) *DBInstance {
	return &DBInstance{
		config:       dbConfig,
		dbConnection: nil,
	}
}

func (inst *DBInstance) Open() error {
	strConnection := fmt.Sprintf(ConnectionString,
		inst.config.Host, inst.config.Port, inst.config.Login,
		security.GetSecureStringContent(&inst.config.Password),
		inst.config.DatabaseName)
	security.DestroySecureString(&inst.config.Password)
	var err error
	inst.dbConnection, err = sql.Open("postgres", strConnection)
	if err != nil {
		return fmt.Errorf("fail to open database connection: %w", err)
	}
	if err := inst.dbConnection.Ping(); err != nil {
		return fmt.Errorf("fail to ping database: %w", err)
	}
	inst.dbConnection.SetConnMaxLifetime(time.Minute * 5)
	inst.dbConnection.SetConnMaxIdleTime(time.Minute * 5)
	inst.dbConnection.SetMaxIdleConns(20)
	inst.dbConnection.SetMaxOpenConns(10)
	return nil
}

func (inst *DBInstance) Close() error {
	err := inst.dbConnection.Close()
	if err != nil {
		return fmt.Errorf("error closing database : %w", err)
	}
	return nil
}
