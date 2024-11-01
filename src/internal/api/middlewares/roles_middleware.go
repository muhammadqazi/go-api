package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/utils"
)

func RolesMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole := c.MustGet("role")

		if !utils.RoleChecker(userRole.(string), allowedRoles) {
			c.AbortWithStatusJSON(403, gin.H{
				"status":  false,
				"message": "You do not have the required permissions for this operation",
			})
			return
		}

		c.Next()
	}
}
