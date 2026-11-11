package env

import "testing"

func TestVarMapCreation(t *testing.T) {
	_, err := CreateVarMap()
	if err != nil {
		t.Error("must not fail")
	}
}
