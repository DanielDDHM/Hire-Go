package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func BindAndValidate(c *gin.Context, model interface{}) bool {
	if err := c.Bind(model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "location": "Bind-validate"})
		return false
	}
	if validationErr := validate.Struct(model); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "location": "validation"})
		return false
	}
	return true
}
