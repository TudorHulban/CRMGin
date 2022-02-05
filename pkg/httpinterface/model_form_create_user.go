package httpinterface

/*
Create user handler.
No cache update, only creation.
*/

import (
	"net/http"
	"strconv"
	"time"

	"github.com/TudorHulban/GinCRM/pkg/logic/authentication"
	"github.com/TudorHulban/GinCRM/pkg/persistence"
	"github.com/gin-gonic/gin"
)

// FormCreateUser Structure used for creating user.
type FormCreateUser struct {
	FieldUserCode string `form:"usercode" validate:"required"`
	FieldPassword string `form:"password" validate:"required"`
}

// ResponseCreateUser Structure used to respond to a create user request.
type ResponseCreateUser struct {
	UserID string `json:"userID"`
}

// Verify with:
// curl -v POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/createuser
func (s *HTTPServer) handlerCreateUser(c *gin.Context) {
	var formData FormCreateUser

	if errValid := BindAndValidate(&formData, c); errValid != nil {
		return
	}

	// create user in RDBMS
	salt := authentication.GenerateSALT()
	hashedPassword, errHash := authentication.HASHPassword(formData.FieldPassword, salt)
	if errHash != nil {
		c.AbortWithError(http.StatusInternalServerError, errHash)
		return
	}

	u := persistence.User{
		SecurityGroupID:   1,
		CreatedAt:         time.Now().Unix(),
		LastUpdateAt:      time.Now().Unix(),
		UserCode:          formData.FieldUserCode,
		PasswordLoginForm: formData.FieldPassword,
		PasswordSALT:      salt,
		PasswordHASH:      string(hashedPassword),
	}

	if errCreate := s.crudLogic.AddUser(&u); errCreate != nil {
		c.AbortWithError(http.StatusInternalServerError, errCreate)
		return
	}

	c.JSON(http.StatusOK, ResponseCreateUser{
		UserID: strconv.FormatUint(u.ID, 10),
	})
}

// prepareLoginRoute Method helps with route preparation.
// Routes need to contain the starting slash ex. /route.
func (s *HTTPServer) prepareCreateUserRoute() []route {
	routeLogin := route{
		Group:    endPointGroupAuthorization,
		Endpoint: endPointCreateUser,
		Method:   http.MethodPost,
		Handler:  s.handlerCreateUser,
	}

	return []route{routeLogin}
}
