package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

/*
	ValidateMongoId is a middleware used on routes ending in /:id
	It basically checks if the :id param is valid,
	otherwise it aborts with an error
*/
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
