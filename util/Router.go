package util

import (
	"github.com/gin-gonic/gin"
)

// RouterContextType contains RouterContext's method or contants
type RouterContextType struct {
	Response customResponse
	Context  *gin.Context
}

type customResponse = func(status int, data gin.H)

// RouterContext provides customized context functions
func RouterContext(context *gin.Context) RouterContextType {
	customResponse := response(context)
	result := RouterContextType{Response: customResponse, Context: context}
	return result
}

func response(context *gin.Context) customResponse {
	return func(status int, response gin.H) {
		result := make(gin.H)

		if status == 0 {
			response["status"] = 200
		} else {
			response["status"] = status
		}

		if response["code"] == nil {
			response["code"] = DefaultCode(response["status"])
		}

		if response["message"] == nil {
			response["message"] = DefaultMessage(response["status"])
		}
		if response["status"].(int) >= 400 {
			response["result"] = false
		} else {
			response["result"] = true
		}

		for key, value := range response {
			result[key] = value
		}

		context.JSON(status, result)
	}
}