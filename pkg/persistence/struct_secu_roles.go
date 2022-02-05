package persistence

// SecurityRight Structure defined for persisting security rights.
type SecurityRight struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

// SecurityRole Structure defined for persisting security roles.
type SecurityRole struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

// SecurityDefRole Structure defined for persisting role definition.
// A role is defined by the security rights it has.
type SecurityDefRole struct {
	RoleID  uint8 `validate:"gte=0"`
	RightID uint8 `validate:"gte=0"`
}

// SecurityProfile Structure defined for persisting security profile.
//
// A profile contains a list of roles for each area it needs to have access to.
type SecurityProfile struct {
	ID          uint8  `gorm:"primaryKey"`
	Description string `validate:"required"`
}

// SecurityDefProfile Structure defined for persisting role definition.
// A role is defined by the security rights it has.
type SecurityDefProfile struct {
	ProfileID uint8 `validate:"gte=0"`
	RoleID    uint8 `validate:"gte=0"`
}
