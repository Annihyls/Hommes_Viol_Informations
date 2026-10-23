package database

import (
	"errors"
	"fmt"
	"hvi_server/configserver"
	"hvi_server/constants"
	"hvi_server/security"
	"hvi_server/utils"
)

type UserType int
const MAX_USERTYPE = 3
const (
	EXTERNAL UserType = iota
	INTERNAL
        POSTGRES
)
var userTypeName = map[UserType] string {
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
                        user_login_env_name = string(cstenv.USER_POSTGRES_LOGIN)
                        user_password_env_name = string(cstenv.USER_POSTGRES_PASSWORD)
                case INTERNAL:
                        user_login_env_name = string(cstenv.USER_INTERNAL_LOGIN)
                        user_password_env_name = string(cstenv.USER_INTERNAL_PASSWORD)
                default:
                        user_login_env_name = string(cstenv.USER_EXTERNAL_LOGIN)
                        user_password_env_name = string(cstenv.USER_EXTERNAL_PASSWORD)
        }
        if utils.IsStringEmpty(cfg.Host) {
                return fmt.Errorf("Database Ip address is not defined")
        }
        if utils.IsStringEmpty(cfg.Port) {
                return fmt.Errorf("Database port is not defined")
        }
        if utils.IsStringEmpty(cfg.DatabaseName) {
                return fmt.Errorf("Database name is not defined")
        }
        if utils.IsStringEmpty(cfg.Login) {
                return fmt.Errorf("%s %s", user_login_env_name, cstenv.MSG_ERROR)
        }
        if security.IsSecureStringEmpty(&cfg.Password) {
                return fmt.Errorf("%s %s", user_password_env_name, cstenv.MSG_ERROR)
        }
        return nil
}

func getLoginFromUserType(userType UserType) string {
        var login    string
        switch userType {
                case POSTGRES:
                login = cstenv.Getenv(cstenv.USER_POSTGRES_LOGIN)
                case INTERNAL:
                login = cstenv.Getenv(cstenv.USER_INTERNAL_LOGIN)
                default:
                login = cstenv.Getenv(cstenv.USER_EXTERNAL_LOGIN)
        }
        return login
}

func getPasswordFromUserType(userType UserType) security.SecureString {
        var password security.SecureString
        switch userType {
                case POSTGRES:
                password = cstenv.Sgetenv(cstenv.USER_POSTGRES_PASSWORD)
                case INTERNAL:
                password = cstenv.Sgetenv(cstenv.USER_INTERNAL_PASSWORD)
                default:
                password = cstenv.Sgetenv(cstenv.USER_EXTERNAL_PASSWORD)
        }
        return password
}

func createConnexionConfigWithUserType(userType UserType) (*ConfigDBConnection, error) {
        dbCoCfg := ConfigDBConnection {
                UserType: userType,
                Host: config.Current["DatabaseHost"],
                Port: config.Current["DatabasePort"],
                DatabaseName: config.Current["DatabaseName"],
                Login: getLoginFromUserType(userType),
                Password: getPasswordFromUserType(userType),
        }
        err := checkConfigConformity(&dbCoCfg)
        if err != nil {
                return nil, fmt.Errorf("Database configuration is not conform: %w", err)
        }
        return &dbCoCfg, nil
}

func initDBConfig() error {
        var errs []error
        for i := range MAX_USERTYPE {
                var err error
                g_instances[i].config, err = createConnexionConfigWithUserType(UserType(i))
                if err != nil {
                        errs = append(errs, fmt.Errorf("%s: %w", userTypeName[UserType(i)], err))
                }
        }
        return errors.Join(errs...)
}

func GetDBConfig(userType UserType) *ConfigDBConnection {
	return g_instances[userType].config;
}
