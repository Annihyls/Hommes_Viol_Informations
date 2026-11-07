package environnement

import (
	"fmt"
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

/*
func Getenv(constenv ConstantEnv) string {
	return os.Getenv(string(constenv))
}

func Sgetenv(constenv ConstantEnv) security.SecureString {
	return security.CreateSecureString(os.Getenv(string(constenv)))
}
*/

type EnvVariable struct {
	securecontent map[string]security.SecureString
}

func NewEnvVariable() (*EnvVariable, error) {
	envVar := &EnvVariable{
		securecontent: make(map[string]security.SecureString),
	}
	sstrEnv := []string{
		string(UserPostgresLogin),
		string(UserInternalLogin),
		string(UserExternalLogin),
		string(HVIServerConfiguration),
		string(UserPostgresPassword),
		string(UserInternalPassword),
		string(UserExternalPassword),
	}
	for _, str := range sstrEnv {
		envSstr := security.CreateSecureString(os.Getenv(str))
		if security.IsSecureStringEmpty(&envSstr) {
			return nil, fmt.Errorf("%s %s", str, ErrorMessage)
		}
		envVar.securecontent[str] = envSstr
	}
	return envVar, nil
}

func (ev EnvVariable) Get(constenv ConstantEnv) security.SecureString {
	return ev.securecontent[string(constenv)]
}
