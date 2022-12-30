package response

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(status bool, message string) gin.H {
	return gin.H{"status": status, "message": message}
}

func JSONResponseWithData(status bool, data interface{}) gin.H {
	return gin.H{"status": status, "data": data}
}
