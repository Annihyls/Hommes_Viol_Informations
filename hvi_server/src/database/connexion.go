// Package database implement connections and
package database

import (
	"fmt"
	"hvi_server/environnement"
	"hvi_server/serverconfig"
)

func Init(serverconf *config.ServerConf, envVar *environnement.EnvVariable) error {
	err := initDBConfig(serverconf, envVar)
	if err != nil {
		return fmt.Errorf("fail to init configuration for: %w", err)
	}
	err = initConnections()
	if err != nil {
		return fmt.Errorf("fail to init connection for: %w", err)
	}
	return nil
}

func Close() error {
	err := closeConnections()
	if err != nil {
		return fmt.Errorf("fail to close: %w", err)
	}
	return nil
}
