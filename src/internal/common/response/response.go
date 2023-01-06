package response

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, code int, status bool, message string) {

	c.JSON(code, gin.H{"status": status, "message": message})
}

func JSONResponseWithData(status bool, data interface{}) gin.H {
	return gin.H{"status": status, "data": data}
}
