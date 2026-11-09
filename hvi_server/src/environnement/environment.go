package env

import "hvi_server/security"

type Key string

const (
	UserPostgresLogin      Key = "USER_POSTGRES_LOGIN"
	UserInternalLogin      Key = "USER_INTERNAL_LOGIN"
	UserExternalLogin      Key = "USER_EXTERNAL_LOGIN"
	HVIServerConfiguration Key = "HVI_SERVER_CONF"
)

type SecureKey string

const (
	UserPostgresPassword SecureKey = "USER_POSTGRES_PASSWORD"
	UserInternalPassword SecureKey = "USER_INTERNAL_PASSWORD"
	UserExternalPassword SecureKey = "USER_EXTERNAL_PASSWORD"
)

const errorMessage = "is not defined in your environment. Have you setup your .env file ?"

type Variables interface {
	Get(key Key) string
	SecureGet(key SecureKey) security.SecureString
}
