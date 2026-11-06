package cstenv

import (
	"hvi_server/security"
	"os"
)

type ConstantEnv string

const (
	UserPostgresLogin      ConstantEnv = "USER_POSTGRES_LOGIN"
	UserPostgresPassword   ConstantEnv = "USER_POSTGRES_PASSWORD"
	UserInternalLogin      ConstantEnv = "USER_INTERNAL_LOGIN"
	UserInternalPassword   ConstantEnv = "USER_INTERNAL_PASSWORD"
	UserExternalLogin      ConstantEnv = "USER_EXTERNAL_LOGIN"
	UserExternalPassword   ConstantEnv = "USER_EXTERNAL_PASSWORD"
	HVIServerConfiguration ConstantEnv = "HVI_SERVER_CONF"
)

const ErrorMessage = "is not defined in your environment. Have you setup your .env file ?"

func Getenv(constenv ConstantEnv) string {
	return os.Getenv(string(constenv))
}

func Sgetenv(constenv ConstantEnv) security.SecureString {
	return security.CreateSecureString(os.Getenv(string(constenv)))
}
