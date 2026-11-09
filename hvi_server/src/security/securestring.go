package security

import (
	"runtime"
	"strings"
)

type SecureString struct {
	content   []byte
	destroyed bool
}

func CreateSecureString(str string) SecureString {
	var content []byte
	if len(strings.TrimSpace(str)) == 0 {
		content = nil
	} else {
		content = []byte(str)
	}
	return SecureString{
		content:   content,
		destroyed: false,
	}
}

func IsSecureStringEmpty(sstr *SecureString) bool {
	if sstr == nil || sstr.destroyed || sstr.content == nil {
		return true
	}
	return false
}

func GetSecureStringContent(sstr *SecureString) []byte {
	if sstr == nil || sstr.destroyed || sstr.content == nil {
		return nil
	}
	return sstr.content
}

func DestroySecureString(sstr *SecureString) {
	if sstr == nil {
		return
	}
	if sstr.content == nil {
		sstr.destroyed = true
		return
	}
	for i := range sstr.content {
		sstr.content[i] = 0
	}
	runtime.KeepAlive(sstr.content)
	sstr.content = nil
	sstr.destroyed = true
}
