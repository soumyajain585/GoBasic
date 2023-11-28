package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	StatusCode int
	Message    string
	// DevMessage error
	Body map[string]interface{}
}

func successResponse(c *gin.Context, Message string, Body map[string]interface{}) {
	response := ResponseBody{
		StatusCode: http.StatusOK,
		Message:    Message,
		Body:       Body,
	}
	c.JSON(http.StatusOK, response)
}

func NoDataFoundResponse(c *gin.Context, Message string) {
	response := ResponseBody{
		StatusCode: http.StatusNoContent,
		Message:    Message,
		Body:       map[string]interface{}{},
	}
	c.JSON(http.StatusOK, response)
}

func InternalServerErrorResponse(c *gin.Context, Error error) {
	errorResponse := ResponseBody{
		StatusCode: 500,
		Message:    "Internal Server Error",
		//DevMessage: Error,
		Body: map[string]interface{}{},
	}
	c.JSON(http.StatusInternalServerError, errorResponse)
}

// func ValidationResponse(c *gin.Context, Message string) {
// 	response := ResponseBody{
// 		StatusCode: http.StatusUnprocessableEntity,
// 		Message:    Message,
// 		Body:       map[string]interface{}{},
// 	}
// 	c.JSON(http.StatusUnprocessableEntity, response)
// }

// func unAuthorizedResponse(c *gin.Context, Message string) {
// 	response := ResponseBody{
// 		StatusCode: http.StatusUnauthorized,
// 		Message:    Message,
// 		Body:       map[string]interface{}{},
// 	}
// 	c.JSON(http.StatusUnauthorized, response)
// }

// func BadRequestResponse(c *gin.Context, Message string) {
// 	response := ResponseBody{
// 		StatusCode: http.StatusBadRequest,
// 		Message:    Message,
// 		Body:       map[string]interface{}{},
// 	}
// 	c.JSON(http.StatusBadRequest, response)
// }

// func SessionExpiredErrorResponse(c *gin.Context, Message string) {
// 	response := ResponseBody{
// 		StatusCode: http.StatusUnauthorized,
// 		Message:    Message,
// 		Body:       map[string]interface{}{},
// 	}
// 	c.JSON(http.StatusUnauthorized, response)
// }

// func RequestBodyLogger(c *gin.Context) string {
// 	requestBody, _ := ioutil.ReadAll(c.Request.Body)
// 	rdr1 := ioutil.NopCloser(bytes.NewBuffer(requestBody))
// 	rdr2 := ioutil.NopCloser(bytes.NewBuffer(requestBody))
// 	c.Request.Body = rdr2
// 	return readBody(rdr1)
// }

// func readBody(reader io.Reader) string {
// 	buf := new(bytes.Buffer)
// 	buf.ReadFrom(reader)
// 	s := buf.String()
// 	return s
// }

// type ResponseBodyNew struct {
// 	StatusCode int
// 	Message    string
// 	ErrorCode  string
// 	Body       map[string]interface{}
// }

// func ErrorCodeResponse(c *gin.Context, Message string, ErrorCode string) {
// 	response := ResponseBodyNew{
// 		StatusCode: http.StatusNoContent,
// 		Message:    Message,
// 		ErrorCode:  ErrorCode,
// 		Body:       map[string]interface{}{},
// 	}
// 	c.JSON(http.StatusOK, response)
// }
