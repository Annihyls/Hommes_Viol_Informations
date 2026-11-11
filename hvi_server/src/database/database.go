// Package database implement connections and
package database

import (
	"errors"
	"fmt"
	"hvi_server/environnement"
	"hvi_server/serverconfig"
	"hvi_server/utils/logger"
)

func createInstances(dbConfigs []DBConfig) ([]Instance, error) {
	instances := make([]Instance, MaxUserType)
	for i := range MaxUserType {
		instances[i] = createDBInstance(dbConfigs[i])
		err := instances[i].Open()
		if err != nil {
			for j := i - 1; i >= 0; j-- {
				errfatal := instances[j].Close()
				if errfatal != nil {
					logger.Stderr.Fatalf(
						"Fail to close connection for %s: %v",
						userTypeName[UserType(j)], errfatal)

				}
			}
			return instances, fmt.Errorf("%s: %w", userTypeName[UserType(i)], err)
		}
	}
	return instances, nil
}

func Init(serverconf config.ServerConf, envVariables env.Variables) ([]Instance, error) {
	dbConfigs, err := initDBConfig(serverconf, envVariables)
	if err != nil {
		return nil, fmt.Errorf("fail to init configuration for: %w", err)
	}
	instances, err := createInstances(dbConfigs)
	if err != nil {
		return nil, fmt.Errorf("fail to init connection for: %w", err)
	}
	return instances, nil
}

func instancesClose(instances []Instance) error {
	var errs []error
	for i := range MaxUserType {
		err := instances[i].Close()
		if err != nil {
			errs = append(errs, fmt.Errorf("%s : %w", userTypeName[UserType(i)], err))
		}
	}
	return errors.Join(errs...)
}

func Close(instances []Instance) error {
	err := instancesClose(instances)
	if err != nil {
		return fmt.Errorf("fail to close: %w", err)
	}
	return nil
}
