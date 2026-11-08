package env

import (
	"fmt"
	"hvi_server/security"
	"hvi_server/utils"
	"os"
)

type Key string

const (
	UserPostgresLogin      Key = "USER_POSTGRES_LOGIN"
	UserInternalLogin      Key = "USER_INTERNAL_LOGIN"
	UserExternalLogin      Key = "USER_EXTERNAL_LOGIN"
	HVIServerConfiguration Key = "HVI_SERVER_CONF"
)

type SecureKey string

const (
	UserPostgresPassword SecureKey = "USER_POSTGRES_PASSWORD"
	UserInternalPassword SecureKey = "USER_INTERNAL_PASSWORD"
	UserExternalPassword SecureKey = "USER_EXTERNAL_PASSWORD"
)

const errorMessage = "is not defined in your environment. Have you setup your .env file ?"

type VarMap struct {
	content       map[string]string
	securecontent map[string]security.SecureString
}

func CreateVarMap() (VarMap, error) {
	envVarMap := VarMap{
		content:       make(map[string]string),
		securecontent: make(map[string]security.SecureString),
	}
	keyEnv := []string{
		string(UserPostgresLogin),
		string(UserInternalLogin),
		string(UserExternalLogin),
		string(HVIServerConfiguration),
	}
	for _, key := range keyEnv {
		value := os.Getenv(key)
		if utils.IsStringEmpty(value) {
			return envVarMap, fmt.Errorf("%s %s", key, errorMessage)
		}
		envVarMap.content[key] = value
	}

	keyEnvSecure := []string{
		string(UserPostgresPassword),
		string(UserInternalPassword),
		string(UserExternalPassword),
	}
	for _, key := range keyEnvSecure {
		secureValue := security.CreateSecureString(os.Getenv(key))
		if security.IsSecureStringEmpty(&secureValue) {
			return envVarMap, fmt.Errorf("%s %s", key, errorMessage)
		}
		envVarMap.securecontent[key] = secureValue
	}
	return envVarMap, nil
}

func (vm VarMap) Get(key Key) string {
	return vm.content[string(key)]
}

func (vm VarMap) SecureGet(key SecureKey) security.SecureString {
	return vm.securecontent[string(key)]
}
