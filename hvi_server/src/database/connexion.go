// Package database implement connections and
package database

import (
	"fmt"
	"hvi_server/environnement"
	"hvi_server/serverconfig"
)

var instances []Instance

func Init(serverconf config.ServerConf, envVariables env.Variables) error {
	dbConfigs, err := initDBConfig(serverconf, envVariables)
	if err != nil {
		return fmt.Errorf("fail to init configuration for: %w", err)
	}
	instances, err = initConnections(dbConfigs)
	if err != nil {
		return fmt.Errorf("fail to init connection for: %w", err)
	}
	return nil
}

func Close() error {
	err := closeConnections(instances)
	if err != nil {
		return fmt.Errorf("fail to close: %w", err)
	}
	return nil
}
