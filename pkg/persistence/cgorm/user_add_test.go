package cgorm_test

import (
	"os"
	"testing"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"
	"github.com/TudorHulban/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddUser(t *testing.T) {
	setup.CleanRDBMS()
	require.Nil(t, cgorm.MigrateDBSchema())

	tt := []struct {
		testName    string
		data        persistence.User
		shouldError bool
	}{
		{testName: "validation issues", data: persistence.User{
			UserCode: "abcd",
		}, shouldError: true},
		{testName: "valid data", data: persistence.User{
			SecurityGroupID: 1,
			UserCode:        "abcd",
		}, shouldError: false},
		{testName: "same user code", data: persistence.User{
			SecurityGroupID: 1,
			UserCode:        "abcd",
		}, shouldError: true},
	}

	userCRUD := cgorm.NewUser(log.New(log.DEBUG, os.Stderr, true))

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			if tc.shouldError {
				assert.Error(t, userCRUD.AddUser(&tc.data))
				return
			}
			assert.Nil(t, userCRUD.AddUser(&tc.data))
		})
	}
}
