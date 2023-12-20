package helpers

import (
	"REST-API/common"

	"github.com/gin-gonic/gin"
)

// helper function to detect current env based upon CLI parameters
func GetCurrentEnv(ENV string) string {
	switch ENV {
	case "dev":
		return common.ENV_LOCAL
	case "local":
		return common.ENV_LOCAL
	case "staging":
		return common.ENV_STAGING
	case "prod":
		return common.ENV_PRODUCTION
	case "production":
		return common.ENV_PRODUCTION
	default:
		return common.ENV_LOCAL
	}
}

// helper function to Generate the error response from contoller
func GenerateErrorResponse(ctx *gin.Context, message string, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"status":  statusCode,
		"error":   err.Error(),
	})
}

// helper function to generate the success response from controller
func GenerateSuccessResponse(ctx *gin.Context, message string, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"status":  statusCode,
		"data":    data,
	})
}
