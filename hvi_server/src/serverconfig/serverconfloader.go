/*
Package config implement loading
*/
package config

import (
	"fmt"
	env "hvi_server/environnement"
	"strings"
)

type Key string

const (
	DatabaseHost Key = "DatabaseHost"
	DatabasePort Key = "DatabasePort"
	DatabaseName Key = "DatabaseName"
	Templates    Key = "PathToTemplateFiles"
	Static       Key = "PathToStaticFiles"
)

type ConfFile interface {
	Open(path string) error
	Close() error
	CanReadLine() bool
	ReadLine() (string, error)
}

type ServerConf struct {
	content map[string]string
}

func (c ServerConf) GetDatabaseHost() string {
	return c.content[string(DatabaseHost)]
}

func (c ServerConf) GetDatabasePort() string {
	return c.content[string(DatabasePort)]
}

func (c ServerConf) GetDatabaseName() string {
	return c.content[string(DatabaseName)]
}

func (c ServerConf) GetTemplateDirPath() string {
	return c.content[string(Templates)]
}

func (c ServerConf) GetStaticDirPath() string {
	return c.content[string(Static)]
}

func Load(confFile ConfFile, envVariables env.Variables) (ServerConf, error) {
	var serverConf ServerConf
	confPath := envVariables.Get(env.HVIServerConfiguration)
	err := confFile.Open(confPath)
	if err != nil {
		return serverConf, fmt.Errorf("unable to open configuration file: %w", err)
	}
	defer confFile.Close()
	serverConf = ServerConf{
		content: make(map[string]string),
	}
	for confFile.CanReadLine() {
		line, err := confFile.ReadLine()
		if err != nil {
			return serverConf, fmt.Errorf("unable to parse file: %w", err)
		}
		trimedLine := strings.TrimSpace(line)
		if trimedLine == "" || strings.HasPrefix(trimedLine, "#") {
			continue
		}
		key, val, ok := strings.Cut(trimedLine, "=")
		if !ok {
			return serverConf, fmt.Errorf("invalid line: %q", trimedLine)
		}
		serverConf.content[strings.TrimSpace(key)] = strings.TrimSpace(val)
	}
	return serverConf, nil
}
