package setup

/*
File contains authentication settings, ex.:
1. salt length
2. minimum password length.
*/

const (
	// LenSALT Is lenght for salt.
	LenSALT = 7
	// HASHCost Holder for bcrypt HASH cost.
	HASHCost = 14
	// LenSessionID Holder for session ID length.
	LenSessionID = 20
)
