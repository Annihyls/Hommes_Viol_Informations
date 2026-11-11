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
	DatabaseHost    Key = "DatabaseHost"
	DatabasePort    Key = "DatabasePort"
	DatabaseName    Key = "DatabaseName"
	TemplateDirPath Key = "PathToTemplateFiles"
	StaticDirPath   Key = "PathToStaticFiles"
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

func (c ServerConf) Get(key Key) string {
	return c.content[string(key)]
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
