package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pchihieuu/url-shortner.git/api/database"
)

func DeleteURL(c *gin.Context) {
	shortID := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	err := r.Del(database.Ctx, shortID).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete shortened link",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"message": "Shortened URL Deleted Successfully",
	})
}