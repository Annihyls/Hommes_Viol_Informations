package env

import (
	"hvi_server/security"
	"testing"
)

func getMockEnv(includingExternalPassword bool) map[string]string {
	mockEnv := make(map[string]string)
	mockEnv[string(UserPostgresLogin)] = "t_pos"
	mockEnv[string(UserInternalLogin)] = "t_int"
	mockEnv[string(UserExternalLogin)] = "t_ext"
	mockEnv[string(HVIServerConfiguration)] = "/path/titi"
	mockEnv[string(UserPostgresPassword)] = "t_pswd1"
	mockEnv[string(UserInternalPassword)] = "t_pswd2"
	if includingExternalPassword {
		mockEnv[string(UserExternalPassword)] = "t_pswd3"
	}
	return mockEnv
}

func TestCreateVarMap(t *testing.T) {
	t.Run("must fail if no variables have been set", func(t *testing.T) {
		_, err := CreateVarMap()
		if err == nil {
			t.Error("must fail")
		}
	})

	t.Run("must fail if not all variables have been set", func(t *testing.T) {
		mockEnv := getMockEnv(false)
		for key, value := range mockEnv {
			t.Setenv(key, value)
			_, err := CreateVarMap()
			if err == nil {
				t.Error("must fail")
			}
		}
	})

	t.Run("must create varmap if all variables have been set", func(t *testing.T) {
		mockEnv := getMockEnv(true)
		for key, value := range mockEnv {
			t.Setenv(key, value)
		}
		_, err := CreateVarMap()
		if err != nil {
			t.Error("must succeed")
		}
	})
}

func TestVarMapGet(t *testing.T) {
	t.Run("must be empty if no key found", func(t *testing.T) {
		varMap, err := CreateVarMap()
		if err == nil {
			t.Error("must fail")
		}
		str := varMap.Get(UserPostgresLogin)
		if str != "" {
			t.Errorf("str = '%s', want ''", str)
		}
	})
	t.Run("must contains the correct value if key exist", func(t *testing.T) {
		mockEnv := getMockEnv(true)
		for key, value := range mockEnv {
			t.Setenv(key, value)
		}
		varMap, err := CreateVarMap()
		if err != nil {
			t.Error("must succeed")
		}
		for _, key := range getKeys() {
			if mockEnv[string(key)] != varMap.Get(key) {
				t.Errorf("key must be defined")
			}
		}
	})
}

func TestVarMapSecureGet(t *testing.T) {
	t.Run("must be empty if no key found", func(t *testing.T) {
		varMap, err := CreateVarMap()
		if err == nil {
			t.Error("must fail")
		}
		sstr := varMap.SecureGet(UserPostgresPassword)
		str := string(security.GetSecureStringContent(&sstr))
		if str != "" {
			t.Errorf("str = '%s', want ''", str)
		}
	})
	t.Run("must contains the correct value if key exist", func(t *testing.T) {
		mockEnv := getMockEnv(true)
		for key, value := range mockEnv {
			t.Setenv(key, value)
		}
		varMap, err := CreateVarMap()
		if err != nil {
			t.Error("must succeed")
		}
		for _, key := range getSecureKeys() {
			sstr := varMap.SecureGet(key)
			str := string(security.GetSecureStringContent(&sstr))
			if mockEnv[string(key)] != str {
				t.Errorf("key must be defined")
			}
		}
	})
}
