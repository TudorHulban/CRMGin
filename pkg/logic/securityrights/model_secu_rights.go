package securityrights

// securityProfiles Map of security profiles.
//
// Map KEY is the security profile.
// Map Value is a slice of user rights in profile.
var securityProfiles map[uint8][]uint

// Init Initializes map.
//
// Could be called for resyncing security roles with persistence.
func Init() error {
	return nil
}
