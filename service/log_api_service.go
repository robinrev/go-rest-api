package service

import (
	"fmt"

	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/repository"
)

func AddNewLogApi(data model.LogApi) {
	err := repository.AddNewLogApi(data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("sukses saving log api")
}
