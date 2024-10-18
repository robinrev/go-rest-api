package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robinrev/go-rest-api/api"
	"github.com/robinrev/go-rest-api/initializer"
	"github.com/robinrev/go-rest-api/model"
)

// list go get
// github.com/gin-gonic/gin => rest api
// gorm.io/gorm => orm
// gorm.io/driver/postgres => postgresql driver
// github.com/joho/godotenv => access .env
// github.com/shopspring/decimal => data type for calculating money

func init() {
	initializer.LoadEnv()
	initializer.DBConnection()
	initializer.Migrate(&model.Company{}, &model.LogActivity{}, &model.LogApi{})
}

func main() {
	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run()
}
