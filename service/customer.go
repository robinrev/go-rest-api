package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/repository"
)

func GetAllCustomer(context *gin.Context) {
	log.Println("get all customer")
	data, err := repository.GetAllCustomer()

	if err != nil {
		log.Println("customer not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, data)
}

func GetCustomerById(context *gin.Context) {
	log.Println("get customer by id")
	id := context.Param("id")
	log.Println("id : ", id)
	data, err := repository.GetCustomerById(id)

	if err != nil {
		log.Println("error fetching customer:", err)
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, data)
}

func AddNewCustomer(context *gin.Context) {
	log.Println("add new customer")
	var data model.Customer
	if err := context.BindJSON(&data); err != nil {
		return
	}

	if err := repository.AddNewCustomer(data); err != nil {
		log.Println("Error inserting customer:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.IndentedJSON(http.StatusOK, data)
}

func UpdateCustomer(context *gin.Context) {
	log.Println("update customer")
	var data model.Customer
	if err := context.BindJSON(&data); err != nil {
		return
	}
	err := repository.UpdateCustomer(data)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
	}
	context.IndentedJSON(http.StatusOK, data)
}
