package cgorm

import (
	"fmt"

	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistenceconn"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// SecurityRR Type concentrates security rights, roles and profiles operations.
type SecurityRR struct {
	l *log.LogInfo
}

// NewSecurityRR Constructor for user security related operations in persistence.
func NewSecurityRR(logger *log.LogInfo) persistence.ISecurityRoles {
	return &SecurityRR{
		l: logger,
	}
}

// GetSecurityRightsForRole Returns a slice with security rights given the role ID.
func (op *SecurityRR) GetSecurityRightsForRole(roleID uint8) ([]uint8, error) {
	go op.l.Debugf("Fetching user rights for role ID:%v", roleID)

	var userRoleDef []persistence.SecurityDefRole
	persistenceconn.GetRDBMSConn().Where("role_id = ?", roleID).Find(&userRoleDef)

	if userRoleDef == nil {
		return nil, fmt.Errorf("no security roles found for role:%v", roleID)
	}

	res := make([]uint8, len(userRoleDef))
	for i, definition := range userRoleDef {
		res[i] = definition.RightID
	}

	go op.l.Debugf("User rights for role ID:%v are:%v", roleID, res)
	return res, nil
}

// GetSecurityRightsForProfile Returns a slice with security rights given the profile ID.
func (op *SecurityRR) GetSecurityRightsForProfile(profileID uint8) ([]uint8, error) {
	go op.l.Debugf("Fetching user rights for profile ID:%v", profileID)

	var userProfilesDef []persistence.SecurityDefProfile
	persistenceconn.GetRDBMSConn().Where("profile_id = ?", profileID).Find(&userProfilesDef)

	userRoles := make([]uint8, len(userProfilesDef))
	for i, definition := range userProfilesDef {
		userRoles[i] = definition.RoleID
	}

	var res []uint8
	for _, roleID := range userRoles {
		buf, errGet := op.GetSecurityRightsForRole(roleID)
		if errGet != nil {
			return nil, errors.WithMessagef(errGet, "when fetching security rights for role ID:%v", roleID)
		}

		go op.l.Debugf("Appending for profile ID:%v following user rights:%v", profileID, buf)
		res = append(res, buf...)
	}

	return res, nil
}

func (op *SecurityRR) GetSecurityProfilesDefinition() (map[uint8][]uint8, error) {
	go op.l.Debug("Fetching security profiles definition:")

	var userProfiles []persistence.SecurityProfile
	persistenceconn.GetRDBMSConn().Find(&userProfiles)

	if len(userProfiles) == 0 {
		return nil, errors.New("did not find any profile")
	}

	profiles := make([]uint8, len(userProfiles))
	for i, profile := range userProfiles {
		profiles[i] = profile.ID
	}

	go op.l.Debugf("Found profiles:%v", profiles)

	res := make(map[uint8][]uint8, len(profiles))

	for profileID := range profiles {
		go op.l.Debugf("Fetching security rights for profile ID:%v", profileID)

		secuRights, errOp := op.GetSecurityRightsForProfile(uint8(profileID))
		if errOp != nil {
			return nil, errors.WithMessagef(errOp, "error when fetching security rights for profile ID:%v", profileID)
		}

		go op.l.Debugf("Security rights for profile ID:%v are:%v", profileID, secuRights)
		res[uint8(profileID)] = secuRights
	}

	return res, nil
}
