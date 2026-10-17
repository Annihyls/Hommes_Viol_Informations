package database

import (
	"hvi_server/constants"
	"hvi_server/utils"
	"hvi_server/security"
	"os"
	"log"
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

func checkIfConfigIsConform(cfg *ConfigDBConnection) {
        const msg_error string = "is not defined in your environment. Have you setup your .env file ?"
        var user_login_env_name, user_password_env_name string
        switch cfg.UserType {
                case POSTGRES:
                user_login_env_name = cstenv.USER_POSTGRES_LOGIN
                user_password_env_name = cstenv.USER_POSTGRES_PASSWORD
                case INTERNAL:
                user_login_env_name = cstenv.USER_INTERNAL_LOGIN
                user_password_env_name = cstenv.USER_INTERNAL_PASSWORD
                default:
                user_login_env_name = cstenv.USER_EXTERNAL_LOGIN
                user_password_env_name = cstenv.USER_EXTERNAL_PASSWORD
        }
        if utils.IsStringEmpty(cfg.Host) {
                log.Fatalf("%s %s", cstenv.DB_HOST, msg_error)
        }
        if utils.IsStringEmpty(cfg.Port) {
                log.Fatalf("%s %s", cstenv.DB_PORT, msg_error)
        }
        if utils.IsStringEmpty(cfg.DatabaseName) {
                log.Fatalf("%s %s", cstenv.DB_NAME, msg_error)
        }
        if utils.IsStringEmpty(cfg.Login) {
                log.Fatalf("%s %s", user_login_env_name, msg_error)
        }
        if security.IsSecureStringEmpty(&cfg.Password) {
                log.Fatalf("%s %s", user_password_env_name, msg_error)
        }
}

func getLoginFromUserType(userType UserType) string {
        var login    string
        switch userType {
                case POSTGRES:
                login = os.Getenv(cstenv.USER_POSTGRES_LOGIN)
                case INTERNAL:
                login = os.Getenv(cstenv.USER_INTERNAL_LOGIN)
                default:
                login = os.Getenv(cstenv.USER_EXTERNAL_LOGIN)
        }
        return login
}

func getPasswordFromUserType(userType UserType) security.SecureString {
        var password security.SecureString
        switch userType {
                case POSTGRES:
                password = security.CreateSecureString(os.Getenv(cstenv.USER_POSTGRES_PASSWORD))
                case INTERNAL:
                password = security.CreateSecureString(os.Getenv(cstenv.USER_INTERNAL_PASSWORD))
                default:
                password = security.CreateSecureString(os.Getenv(cstenv.USER_EXTERNAL_PASSWORD))
        }
        return password
}

func createConnexionConfigWithUserType(userType UserType) ConfigDBConnection {
        dbCoCfg := ConfigDBConnection {
                UserType: userType,
                Host: os.Getenv(cstenv.DB_HOST),
                Port: os.Getenv(cstenv.DB_PORT),
                DatabaseName: os.Getenv(cstenv.DB_NAME),
                Login: getLoginFromUserType(userType),
                Password: getPasswordFromUserType(userType),
        }
        checkIfConfigIsConform(&dbCoCfg)
        return dbCoCfg
}

func initConfig() {
	g_instances[EXTERNAL].config = createConnexionConfigWithUserType(EXTERNAL)
        g_instances[INTERNAL].config = createConnexionConfigWithUserType(INTERNAL)
        g_instances[POSTGRES].config = createConnexionConfigWithUserType(POSTGRES)
}

func GetConfig(userType UserType) ConfigDBConnection {
	return g_instances[userType].config;
}
