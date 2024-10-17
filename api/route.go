package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robinrev/go-rest-api/controller"
	"github.com/robinrev/go-rest-api/service"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(RequestTimingMiddleware())
	routes := router.Group("/companies")
	{
		routes.GET("/", service.GetAllCompany)
		router.GET("/companies/:id", service.GetCompanyById)
		// router.GET("/companies", service.GetCompanyByCompanyCode)
		routes.PATCH("/:companyCode", service.UpdateCompany)
		routes.POST("/", service.AddNewCompany)
	}

	test := router.Group("/test")
	{
		test.GET("/", controller.GetTodos)
	}
}

// CustomResponseWriter wraps the default Gin ResponseWriter to capture the response body
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write method captures the response body in the buffer
func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Capture the response body
	return w.ResponseWriter.Write(b)
}

func RequestTimingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a custom response writer and wrap Gin's ResponseWriter
		customWriter := &CustomResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = customWriter // Replace the context's writer with our custom one

		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		// Restore the body to the request (since reading it consumes the body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// Log the request body
		fmt.Println("Request Body:", string(bodyBytes))

		// Capture and log headers
		for key, values := range c.Request.Header {
			for _, value := range values {
				fmt.Printf("Header: %s = %s\n", key, value)
			}
		}

		startTime := time.Now()
		c.Next()
		// Capture the status, request method, and path
		statusCode := c.Writer.Status() // Get response status code
		requestMethod := c.Request.Method
		requestPath := c.Request.URL.Path

		fmt.Println("Response Body:", customWriter.body.String())
		// Log request and response details after the request has been processed
		log.Printf("Request: %s %s | Status: %d | Time: %v\n",
			requestMethod, requestPath, statusCode, time.Since(startTime))
	}
}
