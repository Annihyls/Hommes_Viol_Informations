package database

import (
	"fmt"
	"hvi_server/constants"
	"hvi_server/security"
	"hvi_server/utils"
)

type UserType int
const (
	EXTERNAL UserType = iota
	INTERNAL
        POSTGRES
)
const MAX_USERTYPE = 3

type ConfigDBConnection struct {
	UserType     UserType
	DatabaseName string
	Host         string
	Port         string
	Login        string
	Password     security.SecureString
}

func checkConfigConformity(cfg *ConfigDBConnection) error {
        const msg_error string = "is not defined in your environment. Have you setup your .env file ?"
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
                return fmt.Errorf("%s %s", cstenv.DB_HOST, msg_error)
        }
        if utils.IsStringEmpty(cfg.Port) {
                return fmt.Errorf("%s %s", cstenv.DB_PORT, msg_error)
        }
        if utils.IsStringEmpty(cfg.DatabaseName) {
                return fmt.Errorf("%s %s", cstenv.DB_NAME, msg_error)
        }
        if utils.IsStringEmpty(cfg.Login) {
                return fmt.Errorf("%s %s", user_login_env_name, msg_error)
        }
        if security.IsSecureStringEmpty(&cfg.Password) {
                return fmt.Errorf("%s %s", user_password_env_name, msg_error)
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
                Host: cstenv.Getenv(cstenv.DB_HOST),
                Port: cstenv.Getenv(cstenv.DB_PORT),
                DatabaseName: cstenv.Getenv(cstenv.DB_NAME),
                Login: getLoginFromUserType(userType),
                Password: getPasswordFromUserType(userType),
        }
        err := checkConfigConformity(&dbCoCfg)
        if err != nil {
                return nil, fmt.Errorf("Database configuration is not conform: %w", err)
        }
        return &dbCoCfg, nil
}

func initConfig() error {
        var err error
	g_instances[EXTERNAL].config, err = createConnexionConfigWithUserType(EXTERNAL)
	if err != nil {
                return fmt.Errorf("External user: %w", err)
	}
	
        g_instances[INTERNAL].config, err = createConnexionConfigWithUserType(INTERNAL)
        if err != nil {
                return fmt.Errorf("Internal user: %w", err)
	}
	
        g_instances[POSTGRES].config, err = createConnexionConfigWithUserType(POSTGRES)
        if err != nil {
                return fmt.Errorf("Postgres user: %w", err)
	}
        return nil
}

func GetConfig(userType UserType) *ConfigDBConnection {
	return g_instances[userType].config;
}
