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
	dbHost := "dbhost"
	dbPort := "dbport"
	dbName := "dbname"
	pathTmpl := "/test/path/templates"
	pathStic := "/test/path2/static"

	rlStrs := []string{
		"#comment",
		"",
		"config=toto",
		string(DatabaseHost) + " = " + dbHost,
		string(DatabasePort) + " = " + dbPort,
		string(DatabaseName) + " = " + dbName,
		string(Templates) + "    = " + pathTmpl,
		string(Static) + "       = " + pathStic,
	}

	m := &MockConfFile{
		ReadLineIteration: 0,
		CanReadLineResult: []bool{true, true, true},
		ReadLineStr:       rlStrs,
	}
	ev := NewMockVarMap()
	ev.content[string(env.HVIServerConfiguration)] = "toto"
	config, err := Load(m, ev)
	if err != nil {
		t.Error("must not fail")
	}

	if config.GetDatabaseHost() == dbHost {
		t.Errorf("must be equal to %s", dbHost)
	}

	if config.GetDatabasePort() == dbPort {
		t.Errorf("must be equal to %s", dbPort)
	}

	if config.GetDatabaseName() == dbName {
		t.Errorf("must be equal to %s", dbName)
	}

	if config.GetTemplateDirPath() == pathTmpl {
		t.Errorf("must be equal to %s", pathTmpl)
	}

	if config.GetStaticDirPath() == pathStic {
		t.Errorf("must be equal to %s", pathStic)
	}
}
