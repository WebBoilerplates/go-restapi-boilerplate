package util

import (
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// RouterContextType contains RouterContext's method or contants
type RouterContextType struct {
	Response  customResponse
	ParseBody func(io.ReadCloser) (map[string]interface{}, error)
	Context   *gin.Context
}

type customResponse = func(status int, data gin.H)

// RouterContext provides customized context functions
func RouterContext(context *gin.Context) RouterContextType {
	customResponse := response(context)
	result := RouterContextType{
		Response:  customResponse,
		Context:   context,
		ParseBody: parseBody,
	}
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

// func errorHandler

// Create error handler, make error type and get by parameters

func parseBody(body io.ReadCloser) (map[string]interface{}, error) {
	bodyValue, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	parseResult, parseErr := ReadJSONByte(bodyValue)
	if parseErr != nil {
		return nil, parseErr
	}

	return parseResult, nil
}
