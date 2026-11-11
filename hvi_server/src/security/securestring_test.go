package security

import "testing"

func TestCreateEmptySecureString(t *testing.T) {

	t.Run("empty string", func(t *testing.T) {
		sstring := CreateSecureString("")
		if sstring.content != nil {
			t.Error("sstring.content != nil; want nil")
		}
		if sstring.destroyed == true {
			t.Error("sstring.destroyed == true; want false")
		}
	})

	t.Run("blank string", func(t *testing.T) {
		sstring := CreateSecureString("        ")
		if sstring.content != nil {
			t.Error("sstring.content != nil; want nil")
		}
		if sstring.destroyed == true {
			t.Error("sstring.destroyed == true; want false")
		}
	})

	t.Run("filled string", func(t *testing.T) {
		sstring := CreateSecureString("titi")
		if sstring.content == nil {
			t.Error("sstring.content == nil; want []byte")
		}
		if sstring.destroyed == true {
			t.Error("sstring.destroyed == true; want false")
		}
	})
}

func TestIsSecureStringEmpty(t *testing.T) {
	t.Run("with nil parameter", func(t *testing.T) {
		wanted := true
		result := IsSecureStringEmpty(nil)
		if result != wanted {
			t.Error("IsSecureStringEmpty(nil) = false; want true")
		}
	})

	t.Run("with nil content", func(t *testing.T) {
		wanted := true
		sstring := SecureString{content: nil, destroyed: false}
		result := IsSecureStringEmpty(&sstring)
		if result != wanted {
			t.Error("IsSecureStringEmpty(&sstring) = false; want true")
		}
	})

	t.Run("with destroyed secure string", func(t *testing.T) {
		wanted := true
		sstring := SecureString{content: []byte{'b', 't', 'p'}, destroyed: true}
		result := IsSecureStringEmpty(&sstring)
		if result != wanted {
			t.Error("IsSecureStringEmpty(&sstring) = false; want true")
		}
	})

	t.Run("with normal secure string", func(t *testing.T) {
		wanted := false
		sstring := SecureString{content: []byte{'b', 't', 'p'}, destroyed: false}
		result := IsSecureStringEmpty(&sstring)
		if result != wanted {
			t.Error("IsSecureStringEmpty(nil) = true; want false")
		}
	})
}

func TestGetSecureStringContent(t *testing.T) {
	t.Run("with nil parameter", func(t *testing.T) {
		result := GetSecureStringContent(nil)
		if result != nil {
			t.Error("GetSecureStringContent(nil) = []byte; want nil")
		}
	})

	t.Run("with nil content", func(t *testing.T) {
		sstring := SecureString{content: nil, destroyed: false}
		result := GetSecureStringContent(&sstring)
		if result != nil {
			t.Error("GetSecureStringContent(&sstring) = []byte; want nil")
		}
	})

	t.Run("with destroyed secure string", func(t *testing.T) {
		sstring := SecureString{content: []byte{'b', 't', 'p'}, destroyed: true}
		result := GetSecureStringContent(&sstring)
		if result != nil {
			t.Error("GetSecureStringContent(&sstring) = []byte; want nil")
		}
	})

	t.Run("normal behavior", func(t *testing.T) {
		sstring := SecureString{content: []byte{'b', 't', 'p'}, destroyed: false}
		result := GetSecureStringContent(&sstring)
		if result == nil {
			t.Error("GetSecureStringContent(&sstring) = nil; want []byte")
		}
	})
}

func TestDestroySecureString(t *testing.T) {
	t.Run("with nil parameter", func(t *testing.T) {
		DestroySecureString(nil)
		/*
		   No errors possible currently, the test
		   run just to check a branch
		*/
	})

	t.Run("with nil content", func(t *testing.T) {
		sstring := SecureString{content: nil, destroyed: false}
		DestroySecureString(&sstring)
		if sstring.content != nil {
			t.Error("sstring.content = []byte; want nil")
		}
		if !sstring.destroyed {
			t.Error("sstring.destroyed = false; want true")
		}
	})

	t.Run("normal behavior", func(t *testing.T) {
		sstring := CreateSecureString("MySecure String")
		DestroySecureString(&sstring)
		if sstring.content != nil {
			t.Error("sstring.content = []byte; want nil")
		}
		if !sstring.destroyed {
			t.Error("sstring.destroyed = false; want true")
		}
	})
}
