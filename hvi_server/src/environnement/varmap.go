package env

import (
	"fmt"
	"hvi_server/security"
	"os"
	"strings"
)

type VarMap struct {
	content       map[string]string
	securecontent map[string]security.SecureString
}

func CreateVarMap() (VarMap, error) {
	envVarMap := VarMap{
		content:       make(map[string]string),
		securecontent: make(map[string]security.SecureString),
	}
	keys := getKeys()
	for _, key := range keys {
		value := os.Getenv(string(key))
		if len(strings.TrimSpace(value)) == 0 {
			return envVarMap, fmt.Errorf("%s %s", key, errorMessage)
		}
		envVarMap.content[string(key)] = value
	}

	sKeys := getSecureKeys()
	for _, key := range sKeys {
		secureValue := security.CreateSecureString(os.Getenv(string(key)))
		if security.IsSecureStringEmpty(&secureValue) {
			return envVarMap, fmt.Errorf("%s %s", key, errorMessage)
		}
		envVarMap.securecontent[string(key)] = secureValue
	}
	return envVarMap, nil
}

func (vm VarMap) Get(key Key) string {
	return vm.content[string(key)]
}

func (vm VarMap) SecureGet(key SecureKey) security.SecureString {
	return vm.securecontent[string(key)]
}
