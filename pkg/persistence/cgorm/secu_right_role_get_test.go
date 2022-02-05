package cgorm_test

import (
	"os"
	"testing"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"
	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSecurityRightsByProfile(t *testing.T) {
	setup.CleanRDBMS()
	require.NoError(t, cgorm.MigrateDBSchema())

	l := log.New(log.DEBUG, os.Stderr, true)
	require.NoError(t, cgorm.PopulateSchemaSecurityRoles(l))

	secu := cgorm.NewSecurityRR(l)

	userRights, errOp := secu.GetSecurityRightsForProfile(0)
	require.NoError(t, errOp)
	assert.Greater(t, len(userRights), 0)

	secuProfiles, errPro := secu.GetSecurityProfilesDefinition()
	require.NoError(t, errPro)
	assert.Greater(t, len(secuProfiles), 0)
}
