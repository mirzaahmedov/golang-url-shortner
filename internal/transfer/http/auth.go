package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirzaahmedov/golang-url-shortner/internal/models"
)

func (r *HTTPRouter) handleAuthRegister(c *gin.Context) {
	data := models.UserRegisterRequest{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
		})
	}

	user := &models.User{
		Fullname: data.Fullname,
		Email:    data.Email,
		Password: data.Password,
	}

	if err := user.EncryptPassword(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
		})
	}

	user, err := r.storage.User().Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

func (r *HTTPRouter) handleAuthLogin(c *gin.Context) {
	data := models.UserLoginRequest{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
		})
	}

	user, err := r.storage.User().GetByEmail(data.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
		})
	}

	if !user.ComparePassword(data.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email or password",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
