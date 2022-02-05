package authentication

/*
File details structure to return as result of authentication operation.
Provides only one method which is for logout.
*/

// UserAuthInfo Concentrates response data for authentication operations.
type UserAuthInfo struct {
	UserRoles []uint // A security profile is a list of roles.
	UserID    uint64 // primary key as per user table.
	UserCode  string
	SessionID string // just generated value
}

// Logout Performs clean up from credentials and session cache
func (a UserAuthInfo) Logout() error {
	return logoutCredentialsCache(a.UserCode)
}
