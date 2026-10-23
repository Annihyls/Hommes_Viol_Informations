package config

import (
	"fmt"
	"hvi_server/constants"
	"hvi_server/utils"
	"os"
	"bufio"
	"strings"
)

var Current map[string]string

func Load() error {
        confPath := cstenv.Getenv(cstenv.HVI_SERVER_CONF)
        if utils.IsStringEmpty(confPath) {
                return fmt.Errorf("%s %s", cstenv.HVI_SERVER_CONF, cstenv.MSG_ERROR);
        }
        
        confFile, err := os.Open(confPath)
        if err != nil {
                return fmt.Errorf("Unable to open configuration file: %w", err)
        }
        defer confFile.Close()
        Current = make(map[string]string)
        scanner := bufio.NewScanner(confFile)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			return fmt.Errorf("Invalid line: %q", line)
		}
		Current[strings.TrimSpace(key)] = strings.TrimSpace(val)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Unable to parse file: %w", err)
	}
	return nil
}
