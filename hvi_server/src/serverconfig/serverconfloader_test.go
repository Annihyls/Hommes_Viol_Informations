package config

import (
	"errors"
	"testing"
)

func TestWithOpenFileError(t *testing.T) {
	m := &mockConfFile{
		OpenFile: nil,
		OpenErr:  errors.New("cannot open file"),
	}
	err := Load(m)
	if err == nil {
		t.Error("must fail")
	}
}

/*
func TestWithCloseFileError(t *testing.T) {
	m := &mockConfFile{
		CloseErr: errors.New("cannot close file"),
	}
	err := Load(m)
	if err == nil {
		t.Error("must fail")
	}
}
*/

func TestWithReadLineError(t *testing.T) {
	m := &mockConfFile{
		ReadLineIteration: 0,
		CanReadLineResult: []bool{true},
		ReadLineStr:       []string{""},
		ReadLineErr:       errors.New("cannot close file"),
	}
	err := Load(m)
	if err == nil {
		t.Error("must fail")
	}
}

func TestWithParsingError(t *testing.T) {
	m := &mockConfFile{
		ReadLineIteration: 0,
		CanReadLineResult: []bool{true},
		ReadLineStr:       []string{"toto tata"},
	}
	err := Load(m)
	if err == nil {
		t.Error("must fail")
	}
}

func TestNormalBehavior(t *testing.T) {
	m := &mockConfFile{
		ReadLineIteration: 0,
		CanReadLineResult: []bool{true, true, true},
		ReadLineStr:       []string{"#comment", "", "config = toto"},
	}
	err := Load(m)
	if err != nil {
		t.Error("must not fail")
	}
}
