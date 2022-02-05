package authentication

import "github.com/pkg/errors"

var (
	// ErrorUnknownCredentials Defined error to be used when bad credentials.
	ErrorUnknownCredentials = errors.New("bad credentials")
)
