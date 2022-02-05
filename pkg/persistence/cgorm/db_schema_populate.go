package cgorm

import (
	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// PopulateSchemaSecurityRoles Inserts security rights and roles.
func PopulateSchemaSecurityRoles(l *log.LogInfo) error {
	var errSecu error // reuse var

	go l.Debug("Populating persistence layer with security rights data:")
	for _, right := range setup.SecuRights {
		if errSecu = AddSecurityRight(right, l); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security right: %v", right)
		}
	}
	go l.Debug("Populated security rights.")

	go l.Debug("Populating persistence layer with security roles data:")
	for _, role := range setup.SecuRoles {
		if errSecu = AddSecurityRole(role, l); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security role: %v", role)
		}
	}
	go l.Debug("Populated security roles.")

	go l.Debug("Populating persistence layer with security roles definition:")
	for roleID, theRights := range setup.RolesDefinition {
		if errSecu = AddSecurityRoleDefinition(uint8(roleID), theRights, l); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security role: %v", roleID)
		}
	}
	go l.Debug("Populated security roles definition.")

	go l.Debug("Populating persistence layer with security profiles definition:")
	for profileID, theRights := range setup.ProfilesDefinition {
		if errSecu = AddSecurityProfileDefinition(uint8(profileID), theRights, l); errSecu != nil {
			return errors.WithMessagef(errSecu, "error adding security profile: %v", profileID)
		}
	}
	go l.Debug("Populated security profiles definition.")

	return nil
}
