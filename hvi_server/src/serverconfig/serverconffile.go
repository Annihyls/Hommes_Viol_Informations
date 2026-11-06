package config

import (
	"bufio"
	"fmt"
	"os"
)

type ConfFile interface {
	Open(path string) (*os.File, error)
	Close(f *os.File) error
	CanReadLine(f *os.File) bool
	ReadLine(f *os.File) (string, error)
}

type ServerConfFile struct {
	scanner *bufio.Scanner
}

func NewServerConfFile() *ServerConfFile {
	return &ServerConfFile{
		scanner: nil,
	}
}

func (cf *ServerConfFile) Open(path string) (*os.File, error) {
	return os.Open(path)
}

func (cf *ServerConfFile) Close(f *os.File) error {
	return f.Close()
}

func (cf *ServerConfFile) CanReadLine(f *os.File) bool {
	if cf.scanner == nil {
		cf.scanner = bufio.NewScanner(f)
	}
	return cf.scanner.Scan()
}

func (cf *ServerConfFile) ReadLine(f *os.File) (string, error) {
	if cf.scanner == nil {
		return "", fmt.Errorf("undefined scanner")
	}
	err := cf.scanner.Err()
	if err != nil {
		return "", err
	}
	return cf.scanner.Text(), nil
}
