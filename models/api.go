package models

import "time"

type RApiItem struct {
	Id                string    `json:"id" gorm:"column:id"`
	Name              string    `json:"name" gorm:"column:name"`
	Path              string    `json:"path" gorm:"column:path"`
	GroupId           int64     `json:"group_id" gorm:"column:group_id"`
	PolicyId          int64     `json:"policy_id" gorm:"column:policy_id"`
	Method            int       `json:"method" gorm:"column:method"`
	RequestHeaders    string    `json:"request_headers" gorm:"column:request_headers"`
	BodyType          int       `json:"body_type" gorm:"column:body_type"`
	BodyRaw           string    `json:"body_raw" gorm:"column:body_raw"`
	BodyJson          string    `json:"body_json" gorm:"column:body_json"`
	BodyForm          string    `json:"body_form" gorm:"column:body_form"`
	ResponseData      string    `json:"response_data" gorm:"column:response_data"`
	Frequency         int       `json:"frequency" gorm:"column:frequency"`
	WebMonitorType    int       `json:"web_monitor_type" gorm:"column:web_monitor_type"`
	ServerMonitorType int       `json:"server_monitor_type" gorm:"column:server_monitor_type"`
	ApiMonitorType    int       `json:"api_monitor_type" gorm:"column:api_monitor_type"`
	WebHttpStatusCode int       `json:"web_http_status_code" gorm:"column:web_http_status_code"`
	ApiHttpStatusCode string    `json:"api_http_status_code" gorm:"column:api_http_status_code"`
	ServerMemory      float64   `json:"server_memory" gorm:"column:server_memory"`
	ServerDisk        float64   `json:"server_disk" gorm:"column:server_disk"`
	ServerCpu         float64   `json:"server_cpu" gorm:"column:server_cpu"`
	CheckSSL          int       `json:"check_ssl" gorm:"column:check_ssl"`
	CheckSSLAdvance   int       `json:"check_ssl_advance" gorm:"column:check_ssl_advance"`
	FailedWaitTimes   int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt         LocalTime `json:"created_at" gorm:"column:created_at"`
	Status            int       `json:"status" gorm:"column:status"`
	BasicUser         string    `json:"basic_user" gorm:"column:basic_user"`
	BasicPassword     string    `json:"basic_password" gorm:"column:basic_password"`
}

func GetAllActiveApiList() []RApiItem {
	var result []RApiItem
	sqlStr := `
	SELECT
	al.id,
	al.path,
	al.method,
	al.request_headers,
	al.body_type,
	al.body_raw,
	al.body_json,
	al.body_form,
	al.response_data,
	al.basic_user,
	al.basic_password,
	mp.frequency,
	mp.api_monitor_type,
	mp.api_http_status_code,
	mp.failed_wait_times
FROM
	api_list AS al
LEFT JOIN monitor_policy AS mp ON mp.id = al.policy_id
WHERE
	al.deleted_at IS NULL
AND al.status = 1`
	DB.Raw(sqlStr).Scan(&result)
	return result
}

type RApiLogItem struct {
	DefaultModel
	ApiId        string    `json:"api_id" gorm:"column:api_id"`
	Status       string    `json:"status" gorm:"column:status"`
	StatusCode   int       `json:"status_code" gorm:"column:status_code"`
	Proto        string    `json:"proto" gorm:"column:proto"`
	Elapsed      int64     `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string    `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int       `json:"check_success" gorm:"check_success"`
	Region       int       `json:"region" gorm:"region"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
}

func GetLatestApiLog(apiId string) RApiLogItem {
	var result RApiLogItem
	sqlStr := `select * from api_log where api_id=@apiId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"apiId": apiId,
	}).Scan(&result)
	return result
}

type ApiLogItem struct {
	DefaultModel
	ApiId        string    `json:"api_id" gorm:"column:api_id"`
	Status       string    `json:"status" gorm:"column:status"`
	StatusCode   int       `json:"status_code" gorm:"column:status_code"`
	Proto        string    `json:"proto" gorm:"column:proto"`
	Elapsed      int64     `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string    `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int       `json:"check_success" gorm:"check_success"`
	Region       int       `json:"region" gorm:"region"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
}

func AddApiLog(apiId string, status string, statusCode int, proto string, elapsed int64, responseData string, checkSuccess int, region int) {
	apiLog := ApiLogItem{
		ApiId:        apiId,
		Status:       status,
		StatusCode:   statusCode,
		Proto:        proto,
		Elapsed:      elapsed,
		ResponseData: responseData,
		CheckSuccess: checkSuccess,
		Region:       region,
	}
	DB.Table("api_log").Create(&apiLog)
}
