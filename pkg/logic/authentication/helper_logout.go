package authentication

import (
	"github.com/TudorHulban/GinCRM/pkg/cache/cachecredentials"
	"github.com/pkg/errors"
)

// logoutCredentialsCache Helper deletes credentials to clean up user access.
func logoutCredentialsCache(usercode string) error {
	i, errCache := cachecredentials.GetCache()
	if errCache != nil {
		return errors.WithMessage(errCache, "credentials not cleaned")
	}

	return i.DeleteKVByK([]byte(usercode))
}
