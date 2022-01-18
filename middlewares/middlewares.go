package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func ValidateMongoId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		idIsValid := bson.IsObjectIdHex(id)
		if !idIsValid {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Id",
			})
			return
		}

		c.Next()
	}
}
