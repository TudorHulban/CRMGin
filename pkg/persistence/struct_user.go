package persistence

/*
File contains structures needed for user management.
*/

// User Consolidates data referring to user authentication.
type User struct {
	ID                uint64 `gorm:"primaryKey"`                // user ID, primary key
	SecurityGroupID   uint8  `validate:"required" gorm:"index"` // security group ID
	CreatedAt         int64  // UNIX time for creation time
	LastLoginAt       int64  // UNIX time for last login
	LastUpdateAt      int64  // UNIX time for last update
	UserCode          string `validate:"required" gorm:"uniqueIndex"` // user code, used for login
	PasswordLoginForm string // password coming from login form
	PasswordSALT      string // should not be sent in JSON, exported for ORM
	PasswordHASH      string // should not be sent in JSON, exported for ORM

}

// Contact Is used when defining a app user.
// The app user should have at least one contact.
type Contact struct {
	ID             uint64 `gorm:"primaryKey"` // contact ID, primary key
	UserID         uint64 `validate:"required"`
	FirstName      string `validate:"required" json:"firstname"`
	LastName       string `validate:"required" json:"lasttname"`
	OfficePhoneNo  string
	MobilePhoneNo  string
	CompanyEmail   string
	WorkEmail      string
	AddressHQ      string
	AddressOffice  string
	AddressBilling string
}

// UserData Consolidates user information for creating a user in persistance layer.
type UserData struct {
	User
	Contacts []Contact
}
