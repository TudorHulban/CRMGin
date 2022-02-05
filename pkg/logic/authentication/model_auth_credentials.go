package authentication

import (
	"github.com/TudorHulban/GinCRM/pkg/cache/cachecredentials"
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/badgerwrap"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// Credentials Auth Concentrates user authentication data based on user and passowrd credentials.
type Credentials struct {
	Code     string
	Password string
}

// OPAuthenticationCredentials Structure defined for the authentication operations.
type OPAuthenticationCredentials struct {
	crudLogic         persistence.IUserCRUD // access CRUD persistance layer
	data              Credentials
	AuthenticatedUser UserAuthInfo // field to be populated in case of succesful login
	l                 *log.LogInfo
	useCache          bool
}

// NewOPAuthenticationCredentials Calling authentication one would need a constructor
// for authentication operation.
func NewOPAuthenticationCredentials(credentials Credentials, crud persistence.IUserCRUD, logger *log.LogInfo) *OPAuthenticationCredentials {
	return &OPAuthenticationCredentials{
		useCache:  true,
		data:      credentials,
		crudLogic: crud,
		l:         logger,
	}
}

// NewOPAuthenticationCredentialsNoCache Calling authentication one would need a constructor
// for authentication operation.
func NewOPAuthenticationCredentialsNoCache(credentials Credentials, crud persistence.IUserCRUD, logger *log.LogInfo) *OPAuthenticationCredentials {
	return &OPAuthenticationCredentials{
		useCache:  false,
		data:      credentials,
		crudLogic: crud,
		l:         logger,
	}
}

// CanLogin Method checks if credentials match the cache or persisted.
func (op *OPAuthenticationCredentials) CanLogin() error {
	// check if credentials exist in cache
	if op.useCache {
		if errCache := op.isCachedAuthenticated(); errCache == nil {
			go op.l.Debug("PASS Authentication based on Credentials CACHE")
			return nil
		}
	}

	// check if credentials persisted
	if errPersisted := op.isPersistentAuthenticated(); errPersisted == nil {
		go op.l.Debug("PASS Authentication based on Credentials Persistance")
		return nil
	}

	return ErrorUnknownCredentials
}

// isCachedAUthenticated Checks if app user is cached authenticated
func (op *OPAuthenticationCredentials) isCachedAuthenticated() error {
	i, errCache := cachecredentials.GetCache()
	if errCache != nil {
		return errors.WithMessage(errCache, "could not access cache")
	}

	pass, errGet := i.GetVByK([]byte(op.data.Code))
	if errGet != nil {
		return errors.WithMessage(errGet, "when checking login cache")
	}

	// TODO: refactor for hashed password
	if op.data.Password != string(pass) {
		return errors.New("password does not match login cache one")
	}

	return nil
}

// isPersistentAuthenticated Checks if user credentails are according to persisted values.
func (op *OPAuthenticationCredentials) isPersistentAuthenticated() error {
	userData, errGet := op.crudLogic.GetUserByCredentials(op.data.Code, op.data.Password)
	if errGet != nil {
		return errGet
	}
	if userData == nil {
		go op.l.Debugf("No user with user/passwd: %v/%v", op.data.Code, op.data.Password)
		return ErrorUnknownCredentials
	}

	if checkPasswordHash(op.data.Password, userData.PasswordSALT, userData.PasswordHASH) {
		// credentials match, fill authentication info
		op.AuthenticatedUser = UserAuthInfo{
			UserID:    userData.ID,
			UserCode:  userData.UserCode,
			SessionID: generateSessionID(),
		}

		go op.l.Debugf("authenticated user: %v", op.AuthenticatedUser)

		// save to credentials cache
		// TODO: place in goroutine
		if errCredenCache := op.saveToCredentialsCache(); errCredenCache != nil {
			go op.l.Debugf("error: %v when saving to credentials cache", errCredenCache)
			return errCredenCache
		}

		return op.saveToSessionCache()
	}

	go op.l.Debugf("Bad password: %v for user: %v", op.data.Password, userData.ID)
	return ErrorUnknownCredentials
}

// saveToLoginCache Method saves credentials to login cache.
func (op *OPAuthenticationCredentials) saveToCredentialsCache() error {
	i, errCache := cachecredentials.GetCache(op.l)
	if errCache != nil {
		return errors.WithMessage(errCache, "could not access cache")
	}

	return i.Set(badgerwrap.KV{
		Key:   []byte(op.data.Code),
		Value: []byte(op.data.Password),
	})
}

// saveToSessionCache Method saves credentials to session ID cache.
// Method placed in this file as operation par of user creation and (re)login.
func (op *OPAuthenticationCredentials) saveToSessionCache() error {
	i, errCache := cachecredentials.GetCache(op.l)
	if errCache != nil {
		return errors.WithMessage(errCache, "could not access cache")
	}

	return i.Set(badgerwrap.KV{
		Key:   []byte(op.data.Code),
		Value: []byte(op.AuthenticatedUser.SessionID),
	})
}
