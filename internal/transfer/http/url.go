package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirzaahmedov/golang-url-shortner/internal/models"
)

func (r *HTTPRouter) handleURLCreate(c *gin.Context) {
	data := &models.URLCreateRequest{}

	if err := c.ShouldBindJSON(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
		})
		return
	}

	url := &models.URL{
		Full: data.Full,
	}

	url, err := r.storage.URL().Create(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": url,
	})
}
