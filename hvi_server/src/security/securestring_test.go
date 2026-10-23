package security

import "testing"

func TestCreateEmptySecureString(t *testing.T) {
        sstring := CreateSecureString("")
        if sstring.content != nil {
                t.Error("sstring.content != nil; want nil")
        }
        if sstring.destroyed == true {
                t.Error("sstring.destroyed == true; want false")
        }
}

func TestCreateSecureString(t *testing.T) {
        sstring := CreateSecureString("titi")
        if sstring.content == nil {
                t.Error("sstring.content == nil; want []byte")
        }
        if sstring.destroyed == true {
                t.Error("sstring.destroyed == true; want false")
        }
}

func TestIsSecureStringEmptyWithNil(t *testing.T) {
        wanted := true
        result := IsSecureStringEmpty(nil)
        if result != wanted {
                t.Error("IsSecureStringEmpty(nil) = false; want true")
        }
}

func TestIsSecureStringEmptyWithNilContent(t *testing.T) {
        wanted := true
        sstring := SecureString { content: nil, destroyed: false }
        result := IsSecureStringEmpty(&sstring)
        if result != wanted {
                t.Error("IsSecureStringEmpty(&sstring) = false; want true")
        }
}

func TestIsSecureStringEmptyWithDestroyedString(t *testing.T) {
        wanted := true
        sstring := SecureString { content: []byte{'b', 't', 'p'}, destroyed: true }
        result := IsSecureStringEmpty(&sstring)
        if result != wanted {
                t.Error("IsSecureStringEmpty(&sstring) = false; want true")
        }
}

func TestIsSecureStringEmptyWithSecureString(t *testing.T) {
        wanted := false
        sstring := SecureString { content: []byte{'b', 't', 'p'}, destroyed: false }
        result := IsSecureStringEmpty(&sstring)
        if result != wanted {
                t.Error("IsSecureStringEmpty(nil) = true; want false")
        }
}

func TestGetSecureStringEmptyWithNil(t *testing.T) {
        result := GetSecureStringContent(nil)
        if result != nil {
                t.Error("GetSecureStringContent(nil) = []byte; want nil")
        }
}

func TestGetSecureStringEmptyWithNilContent(t *testing.T) {
        sstring := SecureString { content: nil, destroyed: false }
        result := GetSecureStringContent(&sstring)
        if result != nil {
                t.Error("GetSecureStringContent(&sstring) = []byte; want nil")
        }
}

func TestGetSecureStringEmptyWithDestroyedString(t *testing.T) {
        sstring := SecureString { content: []byte{'b', 't', 'p'}, destroyed: true }
        result := GetSecureStringContent(&sstring)
        if result != nil {
                t.Error("GetSecureStringContent(&sstring) = []byte; want nil")
        }
}

func TestGetSecureStringEmptyWithSecureString(t *testing.T) {
        sstring := SecureString { content: []byte{'b', 't', 'p'}, destroyed: false }
        result := GetSecureStringContent(&sstring)
        if result == nil {
                t.Error("GetSecureStringContent(&sstring) = nil; want []byte")
        }
}

func TestDestroySecureStringWithNil(t *testing.T) {
        var sstring *SecureString = nil
        DestroySecureString(sstring)
        if sstring != nil {
                t.Error("sstring = SecureString; want nil")
        }
}

func TestDestroySecureStringWithNilContent(t *testing.T) {
        sstring := SecureString { content: nil, destroyed: false }
        DestroySecureString(&sstring)
        if sstring.content != nil {
                t.Error("sstring.content = []byte; want nil")
        }
        if !sstring.destroyed {
                t.Error("sstring.destroyed = false; want true")
        }
}

func TestDestroySecureString(t *testing.T) {
        sstring := CreateSecureString("MySecure String")
        DestroySecureString(&sstring)
        if sstring.content != nil {
                t.Error("sstring.content = []byte; want nil")
        }
        if !sstring.destroyed {
                t.Error("sstring.destroyed = false; want true")
        }
}
