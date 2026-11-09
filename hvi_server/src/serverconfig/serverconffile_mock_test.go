package config

type MockConfFile struct {
	OpenErr error

	CloseErr error

	ReadLineIteration int

	CanReadLineResult []bool

	ReadLineStr []string
	ReadLineErr error
}

func (m *MockConfFile) Open(path string) error {
	return m.OpenErr
}

func (m *MockConfFile) Close() error {
	return m.CloseErr
}

func (m *MockConfFile) CanReadLine() bool {
	result := false
	if m.ReadLineIteration < len(m.CanReadLineResult) {
		result = m.CanReadLineResult[m.ReadLineIteration]
	}
	return result
}

func (m *MockConfFile) ReadLine() (string, error) {
	result := ""
	var err error = nil
	if m.ReadLineIteration < len(m.ReadLineStr) {
		result, err = m.ReadLineStr[m.ReadLineIteration], m.ReadLineErr
		m.ReadLineIteration++
	}
	return result, err
}
