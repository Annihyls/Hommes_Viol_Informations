package cstenv

import (
	"hvi_server/security"
	"os"
)

type ConstantEnv string

const (
        USER_POSTGRES_LOGIN    ConstantEnv = "USER_POSTGRES_LOGIN"
        USER_POSTGRES_PASSWORD ConstantEnv = "USER_POSTGRES_PASSWORD"
        USER_INTERNAL_LOGIN    ConstantEnv = "USER_INTERNAL_LOGIN"
        USER_INTERNAL_PASSWORD ConstantEnv = "USER_INTERNAL_PASSWORD"
        USER_EXTERNAL_LOGIN    ConstantEnv = "USER_EXTERNAL_LOGIN"
        USER_EXTERNAL_PASSWORD ConstantEnv = "USER_EXTERNAL_PASSWORD"
        HVI_SERVER_CONF        ConstantEnv = "HVI_SERVER_CONF"
)

const MSG_ERROR = "is not defined in your environment. Have you setup your .env file ?"
func Getenv(constenv ConstantEnv) string {
        return os.Getenv(string(constenv))
}

func Sgetenv(constenv ConstantEnv) security.SecureString {
        return security.CreateSecureString(os.Getenv(string(constenv)))
}
