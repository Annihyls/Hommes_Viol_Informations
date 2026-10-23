package database

import "fmt"

func Init() error {
        err := initDBConfig()
        if err != nil {
                return fmt.Errorf("Fail to init configuration for: %w", err)
        }
        err = initConnections()
        if err != nil {
                return fmt.Errorf("Fail to init connection for: %w", err)
        }
        return nil
}

func Close() error {
        err := closeConnections()
        if err != nil {
                return fmt.Errorf("Fail to close: %w", err)
        }
        return nil
}
