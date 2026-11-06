package config

import "os"

type mockConfFile struct {
	OpenFile *os.File
	OpenErr  error

	CloseErr error

	ReadLineIteration int

	CanReadLineResult []bool

	ReadLineStr []string
	ReadLineErr error
}

func (m *mockConfFile) Open(path string) (*os.File, error) {
	return m.OpenFile, m.OpenErr
}

func (m *mockConfFile) Close(f *os.File) error {
	return m.CloseErr
}

func (m *mockConfFile) CanReadLine(f *os.File) bool {
	var result bool = false
	if m.ReadLineIteration < len(m.CanReadLineResult) {
		result = m.CanReadLineResult[m.ReadLineIteration]
	}
	return result
}

func (m *mockConfFile) ReadLine(f *os.File) (string, error) {
	var result string = ""
	var err error = nil
	if m.ReadLineIteration < len(m.ReadLineStr) {
		result, err = m.ReadLineStr[m.ReadLineIteration], m.ReadLineErr
		m.ReadLineIteration++
	}
	return result, err
}
