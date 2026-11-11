package config

import (
	"errors"
	env "hvi_server/environnement"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("with open file error", func(t *testing.T) {
		m := &MockConfFile{
			OpenErr: errors.New("cannot open file"),
		}
		ev := NewMockVarMap()
		_, err := Load(m, ev)
		if err == nil {
			t.Error("must fail")
		}
	})

	t.Run("with readline error", func(t *testing.T) {
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
	})

	t.Run("with parsing error", func(t *testing.T) {
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
	})

	t.Run("normal behavior", func(t *testing.T) {
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
			string(TemplateDirPath) + "    = " + pathTmpl,
			string(StaticDirPath) + "       = " + pathStic,
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

		if config.Get(DatabaseHost) == dbHost {
			t.Errorf("must be equal to %s", dbHost)
		}

		if config.Get(DatabasePort) == dbPort {
			t.Errorf("must be equal to %s", dbPort)
		}

		if config.Get(DatabaseName) == dbName {
			t.Errorf("must be equal to %s", dbName)
		}

		if config.Get(TemplateDirPath) == pathTmpl {
			t.Errorf("must be equal to %s", pathTmpl)
		}

		if config.Get(StaticDirPath) == pathStic {
			t.Errorf("must be equal to %s", pathStic)
		}
	})
}
