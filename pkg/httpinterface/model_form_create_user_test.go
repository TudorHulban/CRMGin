package httpinterface

import (
	"net/http"
	"os"
	"testing"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"

	"github.com/TudorHulban/log"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUserFlow(t *testing.T) {
	// clean up
	setup.CleanRDBMS()
	require.NoError(t, cgorm.MigrateDBSchema())

	const socket = "0.0.0.0:8001"
	cfg, _ := CreateConfig(socket, "0.2.0", log.New(log.DEBUG, os.Stderr, true), 1)

	tt := []struct {
		testName       string
		httpMethod     string
		reqURL         string
		usercode       string
		password       string
		statusCodeHTTP int
	}{
		//	{testName: "No Credentials", httpMethod: http.MethodPost, reqURL: endPointGroupAuthorization + endPointLogin, usercode: "", password: "", statusCodeHTTP: http.StatusBadRequest},
		//	{testName: "Bad Credentials", httpMethod: http.MethodPost, reqURL: endPointGroupAuthorization + endPointLogin, usercode: "x", password: "y", statusCodeHTTP: http.StatusUnauthorized},
		{testName: "Create User", httpMethod: http.MethodPost, reqURL: endPointGroupAuthorization + endPointCreateUser, usercode: "john", password: "1234", statusCodeHTTP: http.StatusOK},
		//	{testName: "Login bad pass", httpMethod: http.MethodPost, reqURL: endPointGroupAuthorization + endPointLogin, usercode: "john", password: "12345", statusCodeHTTP: http.StatusUnauthorized},
		{testName: "Login", httpMethod: http.MethodPost, reqURL: endPointGroupAuthorization + endPointLogin, usercode: "john", password: "1234", statusCodeHTTP: http.StatusOK},
	}

	s, errCo := NewGinServer(cfg)
	if assert.Nil(t, errCo) {
		for _, tc := range tt {
			t.Run(tc.testName, func(t *testing.T) {
				apitest.New().
					Handler(s.engine).
					Method(tc.httpMethod).
					URL(tc.reqURL).
					FormData("usercode", tc.usercode).
					FormData("password", tc.password).
					Expect(t).
					Status(tc.statusCodeHTTP).
					End()
			})
		}
	}
}
