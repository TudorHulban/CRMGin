package httpinterface

import (
	"net/http"

	"github.com/TudorHulban/GinCRM/pkg/validator"
	"github.com/gin-gonic/gin"
)

// BindAndValidate Helper takes a pointer and Gin context.
// Returns response to request!
func BindAndValidate(formData interface{}, c *gin.Context) error {
	if errBind := c.Bind(formData); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return errBind
	}

	if errValid := validator.GetValidator().Struct(formData); errValid != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errValid.Error()})
		return errValid
	}

	return nil
}
