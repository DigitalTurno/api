package constants

type contextKey string

const TokenDataKey = contextKey("user")

type TypeToken string

const (
	Token         TypeToken = "TOKEN"
	Refresh       TypeToken = "REFRESH"
	EmailPassword TypeToken = "EMAIL_PASSWORD"
)
