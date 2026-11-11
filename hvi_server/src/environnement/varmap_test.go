package env

import (
	"testing"
)

func TestCreateVarMap(t *testing.T) {
	t.Run("with no environment variable set", func(t *testing.T) {
		_, err := CreateVarMap()
		if err == nil {
			t.Error("must fail")
		}
	})

	t.Run("with partially environment variables set", func(t *testing.T) {
		envMap := map[string]string{
			string(UserPostgresLogin):      "t_pos",
			string(UserInternalLogin):      "t_int",
			string(UserExternalLogin):      "t_ext",
			string(HVIServerConfiguration): "/path/titi",
			string(UserPostgresPassword):   "t_pswd1",
			string(UserInternalPassword):   "t_pswd2",

			// We remove the last because unless,
			// the test will not pass, because
			// all environment variables are defined.
			// string(UserExternalPassword):   "t_pswd3",
		}
		for key, value := range envMap {
			t.Setenv(key, value)
			_, err := CreateVarMap()
			if err == nil {
				t.Error("must fail")
			}
		}
	})

	t.Run("with all environment variables set", func(t *testing.T) {
		envMap := map[string]string{
			string(UserPostgresLogin):      "t_pos",
			string(UserInternalLogin):      "t_int",
			string(UserExternalLogin):      "t_ext",
			string(HVIServerConfiguration): "/path/titi",
			string(UserPostgresPassword):   "t_pswd1",
			string(UserInternalPassword):   "t_pswd2",
			string(UserExternalPassword):   "t_pswd3",
		}
		for key, value := range envMap {
			t.Setenv(key, value)
		}
		_, err := CreateVarMap()
		if err != nil {
			t.Error("must not fail")
		}
	})
}
