package cgorm

import (
	"github.com/pkg/errors"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
	"github.com/TudorHulban/log"
)

// AddSecurityRight Helper for adding security right.
func AddSecurityRight(description string, l *log.LogInfo) error {
	data := persistence.SecurityRight{
		Description: description,
	}
	if errValid := validateStruct(&data); errValid != nil {
		return errors.WithMessage(errValid, "when adding security right")
	}

	return persistenceconn.GetRDBMSConn().Create(&data).Error
}

// AddSecurityRole Helper for adding security role.
func AddSecurityRole(description string, l *log.LogInfo) error {
	data := persistence.SecurityRole{
		Description: description,
	}
	if errValid := validateStruct(&data); errValid != nil {
		return errors.WithMessage(errValid, "when adding security role")
	}

	return persistenceconn.GetRDBMSConn().Create(&data).Error
}

// AddSecurityProfile Helper for adding security role.
func AddSecurityProfile(description string, l *log.LogInfo) error {
	data := persistence.SecurityProfile{
		Description: description,
	}
	if errValid := validateStruct(&data); errValid != nil {
		return errors.WithMessage(errValid, "validation error when adding security profile")
	}

	return persistenceconn.GetRDBMSConn().Create(&data).Error
}

// AddSecurityRoleDefinition Helper for adding security profile.
func AddSecurityRoleDefinition(roleID uint8, roleRights []uint8, l *log.LogInfo) error {
	if roleRights == nil {
		return errors.New("no security rights for passed role")
	}

	data := make([]persistence.SecurityDefRole, len(roleRights))

	for i, rightID := range roleRights {
		data[i] = persistence.SecurityDefRole{
			RoleID:  roleID,
			RightID: rightID,
		}

		if errValid := validateStruct(&data[i]); errValid != nil {
			return errors.WithMessagef(errValid, "when adding security role:%v for security right:%v", roleID, rightID)
		}
	}

	return persistenceconn.GetRDBMSConn().Create(&data).Error
}

// AddSecurityProfileDefinition Helper for adding security profile.
func AddSecurityProfileDefinition(profileID uint8, profileRoles []uint8, l *log.LogInfo) error {
	if profileRoles == nil {
		return errors.New("no roles for passed profile")
	}

	data := make([]persistence.SecurityDefProfile, len(profileRoles))

	for i, roleID := range profileRoles {
		data[i] = persistence.SecurityDefProfile{
			ProfileID: profileID,
			RoleID:    roleID,
		}

		if errValid := validateStruct(&data[i]); errValid != nil {
			return errors.WithMessagef(errValid, "validation error when adding security profile:%v for role:%v", profileID, roleID)
		}
	}

	return persistenceconn.GetRDBMSConn().Create(&data).Error
}
