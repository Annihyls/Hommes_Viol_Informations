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
	Login        security.SecureString
	Password     security.SecureString
}

func getCredentials(envVar *environnement.EnvVariable, userType UserType) (security.SecureString, security.SecureString) {
	var login security.SecureString
	var password security.SecureString
	switch userType {
	case POSTGRES:
		login = envVar.Get(environnement.UserPostgresLogin)
		password = envVar.Get(environnement.UserPostgresPassword)
	case INTERNAL:
		login = envVar.Get(environnement.UserInternalLogin)
		password = envVar.Get(environnement.UserInternalPassword)
	default:
		login = envVar.Get(environnement.UserExternalLogin)
		password = envVar.Get(environnement.UserExternalPassword)
	}
	return login, password
}

func createConnexionConfigWithUserType(serverconf *config.ServerConf, envVar *environnement.EnvVariable, userType UserType) (*ConfigDBConnection, error) {
	login, password := getCredentials(envVar, userType)
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

func initDBConfig(serverconf *config.ServerConf, envVar *environnement.EnvVariable) error {
	var errs []error
	for i := range MaxUserType {
		var err error
		gInstances[i].config, err = createConnexionConfigWithUserType(serverconf, envVar, UserType(i))
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", userTypeName[UserType(i)], err))
		}
	}
	return errors.Join(errs...)
}

func GetDBConfig(userType UserType) *ConfigDBConnection {
	return gInstances[userType].config
}
