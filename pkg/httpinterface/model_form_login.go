package httpinterface

import (
	"net/http"

	"github.com/TudorHulban/GinCRM/pkg/logic/authentication"
	"github.com/gin-gonic/gin"
)

// FormLogin Structure used for validating login request.
type FormLogin struct {
	FieldUserCode string `form:"usercode" validate:"required"`
	FieldPassword string `form:"password" validate:"required"`
}

/*
handlerLogin Handles login with credentials route.
Operations happening:
a. add credentials to user credentials cache
b. generate session ID
c. add session ID to session ID cache
d. return session ID

Verify with:
curl -v POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/login
*/
func (s *HTTPServer) handlerLogin(c *gin.Context) {
	formData := new(FormLogin)

	if errValid := BindAndValidate(formData, c); errValid != nil {
		return
	}

	s.cfg.GLogger.Debug("Login Form Data:", formData)

	// check if authorized. if authorized return session ID.
	// in backend insert in session cache the user structure and in user cache the credentials.
	op := authentication.NewOPAuthenticationCredentialsNoCache(authentication.Credentials{
		Code:     formData.FieldUserCode,
		Password: formData.FieldPassword,
	}, s.crudLogic, s.cfg.GLogger)

	if errAuthenticate := op.CanLogin(); errAuthenticate != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	s.cfg.GLogger.Debug("User Authenticated")
	c.JSON(http.StatusOK, op.AuthenticatedUser.SessionID)
}

// prepareLoginRoute Method helps with route preparation.
// Routes need to contain the starting slash ex. /route.
func (s *HTTPServer) prepareLoginRoute() []route {
	routeLogin := route{
		Group:    endPointGroupAuthorization,
		Endpoint: endPointLogin,
		Method:   http.MethodPost,
		Handler:  s.handlerLogin,
	}

	return []route{routeLogin}
}
