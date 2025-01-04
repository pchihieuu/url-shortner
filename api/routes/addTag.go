package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pchihieuu/url-shortner.git/api/database"
)

type TagRequest struct {
	ShortID     string `json:"shortID"`
	Tag         string `json:"tag"`

}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest
	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	shortId := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortId).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found the given shortID",
		})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		// if the data is not JSON object, assume it as plain string
		data = make(map[string]interface{})
		data["data"] = val
	}

	// check if "tags" fields already exists and it's a slice of strings
	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok{
		for _, t := range existingTags {
            if strTag, ok := t.(string); ok {
                tags = append(tags, strTag)
            }
        }
	}

	// check for duplicate tags
	for _, existingTag := range tags {
		if existingTag == tag {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Tag already exists for the given shortID",
			})
			return
		}
	}

	// Add the new tag to the tags slice
	tags = append(tags, tag)
	data["tags"] = tags

	// Marshal the updated data back to JSON
	updatedData, err := json.Marshal(data)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Marshal updated data",
		})
		return
	}

	err = r.Set(database.Ctx, shortId, updatedData, 0).Err()
	if err !=  nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update the data",
		})
		return
	}
	// Response with the updated data
	c.JSON(http.StatusOK, data)
}