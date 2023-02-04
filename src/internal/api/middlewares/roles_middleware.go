package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/utils"
)

func RolesMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole := c.MustGet("role")

		fmt.Println(userRole)

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
