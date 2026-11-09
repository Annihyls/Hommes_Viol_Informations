package config

import (
	"bufio"
	"fmt"
	"os"
)

type ServerConfFile struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewServerConfFile() *ServerConfFile {
	return &ServerConfFile{
		file:    nil,
		scanner: nil,
	}
}

func (cf *ServerConfFile) Open(path string) error {
	var err error
	cf.file, err = os.Open(path)
	return err
}

func (cf *ServerConfFile) Close() error {
	return cf.file.Close()
}

func (cf *ServerConfFile) CanReadLine() bool {
	if cf.file == nil {
		return false
	}
	if cf.scanner == nil {
		cf.scanner = bufio.NewScanner(cf.file)
	}
	return cf.scanner.Scan()
}

func (cf *ServerConfFile) ReadLine() (string, error) {
	if cf.file == nil {
		return "", fmt.Errorf("undefined file")
	}
	if cf.scanner == nil {
		return "", fmt.Errorf("undefined scanner")
	}
	err := cf.scanner.Err()
	if err != nil {
		return "", err
	}
	return cf.scanner.Text(), nil
}
