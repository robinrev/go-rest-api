package api

import (
	"github.com/gin-gonic/gin"
	"github.com/robinrev/go-rest-api/controller"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(controller.RequestTimingMiddleware())
	companyRoutes := router.Group("/companies")
	{
		companyRoutes.GET("/", controller.GetAllCompanies)
		companyRoutes.GET("/:id", controller.GetCompanyById)
		companyRoutes.PATCH("/", controller.UpdateCompany)
		companyRoutes.POST("/", controller.AddNewCompany)
	}

	customerRoutes := router.Group("/customers")
	{
		customerRoutes.GET("/", controller.GetAllCustomer)
		customerRoutes.GET("/company/:companyId", controller.GetAllCustomersByCompany)
		customerRoutes.GET("/:id", controller.GetCustomerById)
		customerRoutes.POST("/", controller.AddNewCustomer)
		customerRoutes.PATCH("/", controller.UpdateCustomer)
	}
}
