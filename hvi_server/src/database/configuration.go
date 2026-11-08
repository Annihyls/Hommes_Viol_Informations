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

type ConfigDBConnection struct {
	UserType     UserType
	DatabaseName string
	Host         string
	Port         string
	Login        string
	Password     security.SecureString
}

func getCredentials(envVarMap *env.VarMap, userType UserType) (string, security.SecureString) {
	var login string
	var password security.SecureString
	switch userType {
	case POSTGRES:
		login = envVarMap.Get(env.UserPostgresLogin)
		password = envVarMap.SecureGet(env.UserPostgresPassword)
	case INTERNAL:
		login = envVarMap.Get(env.UserInternalLogin)
		password = envVarMap.SecureGet(env.UserInternalPassword)
	default:
		login = envVarMap.Get(env.UserExternalLogin)
		password = envVarMap.SecureGet(env.UserExternalPassword)
	}
	return login, password
}

func createConnexionConfigWithUserType(serverconf *config.ServerConf, envVarMap *env.VarMap, userType UserType) ConfigDBConnection {
	login, password := getCredentials(envVarMap, userType)
	dbCoCfg := ConfigDBConnection{
		UserType:     userType,
		Host:         serverconf.Get("DatabaseHost"),
		Port:         serverconf.Get("DatabasePort"),
		DatabaseName: serverconf.Get("DatabaseName"),
		Login:        login,
		Password:     password,
	}
	return dbCoCfg
}

func initDBConfig(serverconf *config.ServerConf, envVarMap *env.VarMap) ([]ConfigDBConnection, error) {
	var errs []error
	var dbConfigs []ConfigDBConnection
	for i := range MaxUserType {
		dbConfig := createConnexionConfigWithUserType(serverconf, envVarMap, UserType(i))
		dbConfigs = append(dbConfigs, dbConfig)
	}
	return dbConfigs, errors.Join(errs...)
}
