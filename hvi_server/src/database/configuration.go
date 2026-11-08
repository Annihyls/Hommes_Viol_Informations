package database

import (
	"errors"
	"fmt"
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

func createConnexionConfigWithUserType(serverconf *config.ServerConf, envVarMap *env.VarMap, userType UserType) (*ConfigDBConnection, error) {
	login, password := getCredentials(envVarMap, userType)
	dbCoCfg := ConfigDBConnection{
		UserType:     userType,
		Host:         serverconf.Get("DatabaseHost"),
		Port:         serverconf.Get("DatabasePort"),
		DatabaseName: serverconf.Get("DatabaseName"),
		Login:        login,
		Password:     password,
	}
	return &dbCoCfg, nil
}

func initDBConfig(serverconf *config.ServerConf, envVarMap *env.VarMap) error {
	var errs []error
	for i := range MaxUserType {
		var err error
		gInstances[i].config, err = createConnexionConfigWithUserType(serverconf, envVarMap, UserType(i))
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", userTypeName[UserType(i)], err))
		}
	}
	return errors.Join(errs...)
}

func GetDBConfig(userType UserType) *ConfigDBConnection {
	return gInstances[userType].config
}
