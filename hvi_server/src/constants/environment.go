package cstenv

import (
	"hvi_server/security"
	"os"
)

type ConstantEnv string

const (
        DB_HOST                ConstantEnv = "DATABASE_HOST"
        DB_PORT                ConstantEnv = "DATABASE_PORT"
        DB_NAME                ConstantEnv = "DATABASE_NAME"
        USER_POSTGRES_LOGIN    ConstantEnv = "USER_POSTGRES_LOGIN"
        USER_POSTGRES_PASSWORD ConstantEnv = "USER_POSTGRES_PASSWORD"
        USER_INTERNAL_LOGIN    ConstantEnv = "USER_INTERNAL_LOGIN"
        USER_INTERNAL_PASSWORD ConstantEnv = "USER_INTERNAL_PASSWORD"
        USER_EXTERNAL_LOGIN    ConstantEnv = "USER_EXTERNAL_LOGIN"
        USER_EXTERNAL_PASSWORD ConstantEnv = "USER_EXTERNAL_PASSWORD"
        PATH_STATIC            ConstantEnv = "PATH_STATIC"
        PATH_TEMPLATES         ConstantEnv = "PATH_TEMPLATES"
)

func Getenv(constenv ConstantEnv) string {
        return os.Getenv(string(constenv))
}

func Sgetenv(constenv ConstantEnv) security.SecureString {
        return security.CreateSecureString(os.Getenv(string(constenv)))
}
