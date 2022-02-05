package persistence

// ISecurityRoles Methods used in user security rights and roles.
type ISecurityRoles interface {
	GetSecurityRightsForProfile(profileID uint8) ([]uint8, error)
	// Returns map.
	// Map KEY: profile ID.
	// Map Value: slice of user rights.
	GetSecurityProfilesDefinition() (map[uint8][]uint8, error)
}

// IUserCreation Methods used in user creation
type IUserCreation interface {
	// Create User, passing pointer to return created ID
	AddUser(*User) error
}

// IUserRetrieval Methods used in user retrieval
type IUserRetrieval interface {
	GetUserByCredentials(userCode, password string) (*User, error)
}

// IUserCRUD Full specifications for a user creation process
type IUserCRUD interface {
	IUserCreation
	IUserRetrieval
	// Update User
	// Delete USer
}
