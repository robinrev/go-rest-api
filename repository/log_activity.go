package repository

import (
	"time"

	"github.com/robinrev/go-rest-api/initializer"
	"github.com/robinrev/go-rest-api/model"
)

func AddNewLogActivity(activity string) error {
	var logActivity model.LogActivity
	logActivity.LogActivityName = activity
	logActivity.CompanyID = 1
	logActivity.EmployeeID = 1
	logActivity.CreateDate = time.Now()
	logActivity.UpdateDate = time.Now()
	logActivity.RecordDate = time.Now()
	result := initializer.DB.Create(&logActivity)
	return result.Error
}
