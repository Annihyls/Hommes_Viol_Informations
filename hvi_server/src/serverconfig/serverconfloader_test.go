package config

import (
	"errors"
	env "hvi_server/environnement"
	"testing"
)

func TestWithOpenFileError(t *testing.T) {
	m := &MockConfFile{
		OpenErr: errors.New("cannot open file"),
	}
	ev := NewMockVarMap()
	_, err := Load(m, ev)
	if err == nil {
		t.Error("must fail")
	}
}

func TestWithReadLineError(t *testing.T) {
	m := &MockConfFile{
		ReadLineIteration: 0,
		CanReadLineResult: []bool{true},
		ReadLineStr:       []string{""},
		ReadLineErr:       errors.New("cannot close file"),
	}
	ev := NewMockVarMap()
	_, err := Load(m, ev)
	if err == nil {
		t.Error("must fail")
	}
}

func TestWithParsingError(t *testing.T) {
	m := &MockConfFile{
		ReadLineIteration: 0,
		CanReadLineResult: []bool{true},
		ReadLineStr:       []string{"toto tata"},
	}
	ev := NewMockVarMap()
	_, err := Load(m, ev)
	if err == nil {
		t.Error("must fail")
	}
}

func TestNormalBehavior(t *testing.T) {
	pathTmpl := "/test/path/templates"
	pathStic := "/test/path2/static"
	m := &MockConfFile{
		ReadLineIteration: 0,
		CanReadLineResult: []bool{true, true, true},
		ReadLineStr: []string{"#comment",
			"",
			"config = toto",
			string(Templates) + " = " + pathTmpl,
			string(Static) + " = " + pathStic,
		},
	}
	ev := NewMockVarMap()
	ev.content[string(env.HVIServerConfiguration)] = "toto"
	config, err := Load(m, ev)
	if err != nil {
		t.Error("must not fail")
	}

	if config.GetTemplateDirPath() == pathTmpl {
		t.Error("must be equal")
	}

	if config.GetStaticDirPath() == pathStic {
		t.Error("must be equal")
	}
}
