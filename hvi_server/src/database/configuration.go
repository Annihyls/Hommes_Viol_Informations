package database

import (
	"errors"
	"hvi_server/environnement"
	"hvi_server/security"
	"hvi_server/serverconfig"
)

type UserType int

const MaxUserType = 3
const (
	EXTERNAL UserType = iota
	INTERNAL
	POSTGRES
)

var userTypeName = map[UserType]string{
	EXTERNAL: "External user",
	INTERNAL: "Internal user",
	POSTGRES: "Postgres user",
}

type DBConfig struct {
	UserType     UserType
	DatabaseName string
	Host         string
	Port         string
	Login        string
	Password     security.SecureString
}

func getCredentials(envVariables env.Variables, userType UserType) (string, security.SecureString) {
	var login string
	var password security.SecureString
	switch userType {
	case POSTGRES:
		login = envVariables.Get(env.UserPostgresLogin)
		password = envVariables.SecureGet(env.UserPostgresPassword)
	case INTERNAL:
		login = envVariables.Get(env.UserInternalLogin)
		password = envVariables.SecureGet(env.UserInternalPassword)
	default:
		login = envVariables.Get(env.UserExternalLogin)
		password = envVariables.SecureGet(env.UserExternalPassword)
	}
	return login, password
}

func createConnexionConfigWithUserType(serverconf config.ServerConf, envVariables env.Variables, userType UserType) DBConfig {
	login, password := getCredentials(envVariables, userType)
	dbCoCfg := DBConfig{
		UserType:     userType,
		Host:         serverconf.GetDatabaseHost(),
		Port:         serverconf.GetDatabasePort(),
		DatabaseName: serverconf.GetDatabaseName(),
		Login:        login,
		Password:     password,
	}
	return dbCoCfg
}

func initDBConfig(serverconf config.ServerConf, envVariables env.Variables) ([]DBConfig, error) {
	var errs []error
	var dbConfigs []DBConfig
	for i := range MaxUserType {
		dbConfig := createConnexionConfigWithUserType(serverconf, envVariables, UserType(i))
		dbConfigs = append(dbConfigs, dbConfig)
	}
	return dbConfigs, errors.Join(errs...)
}
