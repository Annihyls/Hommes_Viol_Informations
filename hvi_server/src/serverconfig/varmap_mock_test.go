package config

import (
	env "hvi_server/environnement"
	"hvi_server/security"
)

type MockVarMap struct {
	content       map[string]string
	secureContent map[string]security.SecureString
}

func NewMockVarMap() *MockVarMap {
	return &MockVarMap{
		content:       make(map[string]string),
		secureContent: make(map[string]security.SecureString),
	}
}

func (mvm MockVarMap) Get(key env.Key) string {
	return mvm.content[string(key)]
}

func (mvm MockVarMap) SecureGet(key env.SecureKey) security.SecureString {
	return mvm.secureContent[string(key)]
}
