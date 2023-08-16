package responses

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, statusCode int, data interface{}) {
	response := struct {
		Status bool        `json:"status"`
		Data   interface{} `json:"data"`
	}{
		Status: true,
		Data:   data,
	}

	c.AbortWithStatusJSON(statusCode, response)
}

func Error(c *gin.Context, statusCode int, err error) {
	response := struct {
		Status bool   `json:"status"`
		Error  string `json:"message"`
	}{
		Status: false,
		Error:  err.Error(),
	}

	c.AbortWithStatusJSON(statusCode, response)
}
