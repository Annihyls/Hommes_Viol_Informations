package database

import (
	"errors"
	"fmt"
	"hvi_server/constants"
	"hvi_server/security"
	"hvi_server/serverconfig"
	"hvi_server/utils"
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

func checkConfigConformity(cfg *ConfigDBConnection) error {
	var user_login_env_name, user_password_env_name string
	switch cfg.UserType {
	case POSTGRES:
		user_login_env_name = string(cstenv.UserPostgresLogin)
		user_password_env_name = string(cstenv.UserPostgresPassword)
	case INTERNAL:
		user_login_env_name = string(cstenv.UserInternalLogin)
		user_password_env_name = string(cstenv.UserInternalPassword)
	default:
		user_login_env_name = string(cstenv.UserExternalLogin)
		user_password_env_name = string(cstenv.UserExternalPassword)
	}
	if utils.IsStringEmpty(cfg.Host) {
		return fmt.Errorf("database Ip address is not defined")
	}
	if utils.IsStringEmpty(cfg.Port) {
		return fmt.Errorf("database port is not defined")
	}
	if utils.IsStringEmpty(cfg.DatabaseName) {
		return fmt.Errorf("database name is not defined")
	}
	if utils.IsStringEmpty(cfg.Login) {
		return fmt.Errorf("%s %s", user_login_env_name, cstenv.ErrorMessage)
	}
	if security.IsSecureStringEmpty(&cfg.Password) {
		return fmt.Errorf("%s %s", user_password_env_name, cstenv.ErrorMessage)
	}
	return nil
}

func getLoginFromUserType(userType UserType) string {
	var login string
	switch userType {
	case POSTGRES:
		login = cstenv.Getenv(cstenv.UserPostgresLogin)
	case INTERNAL:
		login = cstenv.Getenv(cstenv.UserInternalLogin)
	default:
		login = cstenv.Getenv(cstenv.UserExternalLogin)
	}
	return login
}

func getPasswordFromUserType(userType UserType) security.SecureString {
	var password security.SecureString
	switch userType {
	case POSTGRES:
		password = cstenv.Sgetenv(cstenv.UserPostgresPassword)
	case INTERNAL:
		password = cstenv.Sgetenv(cstenv.UserInternalPassword)
	default:
		password = cstenv.Sgetenv(cstenv.UserExternalPassword)
	}
	return password
}

func createConnexionConfigWithUserType(serverconf *config.ServerConf, userType UserType) (*ConfigDBConnection, error) {
	dbCoCfg := ConfigDBConnection{
		UserType:     userType,
		Host:         serverconf.Get("DatabaseHost"),
		Port:         serverconf.Get("DatabasePort"),
		DatabaseName: serverconf.Get("DatabaseName"),
		Login:        getLoginFromUserType(userType),
		Password:     getPasswordFromUserType(userType),
	}
	err := checkConfigConformity(&dbCoCfg)
	if err != nil {
		return nil, fmt.Errorf("database configuration is not conform: %w", err)
	}
	return &dbCoCfg, nil
}

func initDBConfig(serverconf *config.ServerConf) error {
	var errs []error
	for i := range MaxUserType {
		var err error
		gInstances[i].config, err = createConnexionConfigWithUserType(serverconf, UserType(i))
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", userTypeName[UserType(i)], err))
		}
	}
	return errors.Join(errs...)
}

func GetDBConfig(userType UserType) *ConfigDBConnection {
	return gInstances[userType].config
}
