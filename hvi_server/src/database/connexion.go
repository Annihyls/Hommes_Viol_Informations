// Package database implement connections and
package database

import (
	"fmt"
	"hvi_server/serverconfig"
)

func Init(serverconf *config.ServerConf) error {
	err := initDBConfig(serverconf)
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
