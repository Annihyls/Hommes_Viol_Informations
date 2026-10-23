package utils

import "testing"

func TestWithEmptyString(t *testing.T) {
        wanted := true
        result := IsStringEmpty("")
        if result != wanted {
                t.Error("IsStringEmpty(\"\") = false; want true")
        }
}

func TestWithSpacedString(t *testing.T) {
        wanted := true
        result := IsStringEmpty("       ")
        if result != wanted {
                t.Error("IsStringEmpty(\"       \") = false; want true")
        }
}

func TestWithString(t *testing.T) {
        wanted := false
        result := IsStringEmpty("toto")
        if result != wanted {
                t.Error("IsStringEmpty(\"toto\") = true; want false")
        }
}
