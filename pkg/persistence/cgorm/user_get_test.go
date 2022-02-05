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

func TestGetUserByCredentials(t *testing.T) {
	setup.CleanRDBMS()
	require.NoError(t, cgorm.MigrateDBSchema())

	tt := []struct {
		testName    string
		usercode    string
		password    string
		shouldError bool
	}{
		{testName: "should pass", usercode: "john", password: "1234", shouldError: true},
	}

	userCRUD := cgorm.NewUser(log.New(log.DEBUG, os.Stderr, true))

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			_, errGet := userCRUD.GetUserByCredentials(tc.usercode, tc.password)
			if tc.shouldError {
				assert.Error(t, errGet)
				return
			}
			assert.NoError(t, errGet)
		})
	}
}
