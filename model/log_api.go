package model

type LogApi struct {
	LogApiID     int    `gorm:"column:log_api_id;primaryKey;autoIncrement" json:"logApiId"`
	LogApiName   string `gorm:"column:log_api_name;type:varchar(50)" json:"logApiName"`
	Reference    string `gorm:"column:reference;type:varchar(50)" json:"reference"`
	CreateBy     string `gorm:"column:create_by;type:varchar(50)" json:"createBy"`
	ServiceFrom  string `gorm:"column:service_from;type:varchar(50)" json:"serviceFrom"`
	ServiceTo    string `gorm:"column:service_to;type:varchar(50)" json:"serviceTo"`
	Url          string `gorm:"column:url" json:"url"`
	Method       string `gorm:"column:method;type:varchar(10)" json:"method"`
	Header       string `gorm:"column:header" json:"header"`
	RequestBody  string `gorm:"column:request_body" json:"requestBody"`
	ResponseBody string `gorm:"column:response_body" json:"responseBody"`
	HttpResponse string `gorm:"column:http_response;type:varchar(10)" json:"httpResponse"`
	GeneralTime
}

func (LogApi) TableName() string {
	return "log_api"
}
