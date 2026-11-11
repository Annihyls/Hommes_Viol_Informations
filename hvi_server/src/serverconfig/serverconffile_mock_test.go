package config

type MockConfFile struct {
	OpenErr error

	CloseErr error

	ReadLineIteration int

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
	return m.ReadLineIteration < len(m.ReadLineStr)
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
