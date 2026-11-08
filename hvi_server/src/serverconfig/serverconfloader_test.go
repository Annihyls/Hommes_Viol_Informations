package config

import (
	"errors"
	env "hvi_server/environnement"
	"testing"
)

func TestWithOpenFileError(t *testing.T) {
	m := &mockConfFile{
		OpenFile: nil,
		OpenErr:  errors.New("cannot open file"),
	}
	ev := env.VarMap{}
	_, err := Load(m, ev)
	if err == nil {
		t.Error("must fail")
	}
}

func TestWithReadLineError(t *testing.T) {
	m := &mockConfFile{
		ReadLineIteration: 0,
		CanReadLineResult: []bool{true},
		ReadLineStr:       []string{""},
		ReadLineErr:       errors.New("cannot close file"),
	}
	ev := env.VarMap{}
	_, err := Load(m, ev)
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
	ev := env.VarMap{}
	_, err := Load(m, ev)
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
	ev := env.VarMap{}
	_, err := Load(m, ev)
	if err != nil {
		t.Error("must not fail")
	}
}
