package models

type RHeartbeatItem struct {
	Id              string    `json:"id" gorm:"column:id"`
	Name            string    `json:"name" gorm:"column:name"`
	Token           string    `json:"token" gorm:"column:token"`
	GroupId         int64     `json:"group_id" gorm:"column:group_id"`
	Frequency       int       `json:"frequency" gorm:"column:frequency"`
	FailedWaitTimes int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt       LocalTime `json:"created_at" gorm:"column:created_at"`
	Status          int       `json:"status" gorm:"column:status"`
}

func GetAllActiveHeartbeatList() []RHeartbeatItem {
	var result []RHeartbeatItem
	sqlStr := `
	SELECT
	d.id,
	d.name,
	d.group_id,
	d.status,
	d.created_at,
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
	heartbeat_list as d
LEFT JOIN monitor_policy as mp ON mp.id = d.policy_id
WHERE
	d.deleted_at IS NULL
AND d.status = 1`
	DB.Raw(sqlStr).Scan(&result)
	return result
}
