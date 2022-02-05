package setup

/*
File contains security riths and roles.
*/

// SecuRights Temporary storing security rights.
var SecuRights = []string{
	"Login",
	"Add user",
	"View User",
	"Update user",
	"Set Minimum Password Length",
}

// SecuRoles Temporary storing security roles.
var SecuRoles = []string{
	"User Management",
	"User View Only",
	"Manage Application Settings",
}

// RolesDefinition Temporary storing application security roles definition.
// A role is a slice of user rights IDs.
var RolesDefinition = [][]uint8{
	[]uint8{0, 1, 2},
	[]uint8{0, 2},
	[]uint8{3},
}

// SecuProfiles Temporary storing security profiles.
var SecuProfiles = []string{
	"Administrator",
	"Guest",
}

// ProfilesDefinition Temporary storing application security profiles definition.
// A profile is a slice of user roles IDs.
var ProfilesDefinition = [][]uint8{
	[]uint8{0, 2},
	[]uint8{1},
}
