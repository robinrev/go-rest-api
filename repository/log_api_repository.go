package repository

import (
	"time"

	"github.com/robinrev/go-rest-api/initializer"
	"github.com/robinrev/go-rest-api/model"
)

func AddNewLogApi(logApi model.LogApi) error {
	logApi.RecordDate = time.Now()
	logApi.CreateDate = time.Now()
	logApi.UpdateDate = time.Now()
	result := initializer.DB.Create(&logApi)
	return result.Error
}
