package models

type RWebItem struct {
	Id                string    `json:"id" gorm:"column:id"`
	Name              string    `json:"name" gorm:"column:name"`
	Path              string    `json:"path" gorm:"column:path"`
	GroupId           int64     `json:"group_id" gorm:"column:group_id"`
	Frequency         int       `json:"frequency" gorm:"column:frequency"`
	WebMonitorType    int       `json:"web_monitor_type" gorm:"column:web_monitor_type"`
	ServerMonitorType int       `json:"server_monitor_type" gorm:"column:server_monitor_type"`
	ApiMonitorType    int       `json:"api_monitor_type" gorm:"column:api_monitor_type"`
	WebHttpStatusCode int       `json:"web_http_status_code" gorm:"column:web_http_status_code"`
	WebHttpRegexpText string    `json:"web_http_regexp_text" gorm:"column:web_http_regexp_text"`
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

func GetAllActiveWebList() []RWebItem {
	var result []RWebItem
	sqlStr := `
	SELECT
	d.id,
	d.name,
	d.path,
	d.group_id,
	d.status,
	d.created_at,
	d.basic_user,
	d.basic_password,
	mp.frequency,
	mp.web_monitor_type,
	mp.server_monitor_type,
	mp.api_monitor_type,
	mp.web_http_status_code,
	mp.api_http_status_code,
	mp.server_memory,
	mp.server_disk,
	mp.server_cpu,
	mp.check_ssl,
	mp.check_ssl_advance,
	mp.failed_wait_times
FROM
	web_list AS d
	LEFT JOIN monitor_policy AS mp ON mp.id = d.policy_id 
WHERE
	d.deleted_at IS NULL 
	AND d.status = 1`
	DB.Raw(sqlStr).Scan(&result)
	return result
}
