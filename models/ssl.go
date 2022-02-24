package models

import "time"

type RSSLItem struct {
	TEndTime time.Time `json:"t_end_time" gorm:"column:end_time"`
}

func GetSslConfig(id int64) RSSLItem {
	var result RSSLItem
	sqlStr := `select t_end_time from ssl_config where id=@id and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"id": id,
	}).Scan(&result)
	return result
}

type RActiveSSLItem struct {
	Id              int64  `json:"id" gorm:"column:id"`
	WebId           int64  `json:"web_id" gorm:"column:web_id"`
	Name            string `json:"name" gorm:"column:name"`
	Path            string `json:"path" gorm:"column:path"`
	GroupId         int64  `json:"group_id" gorm:"column:group_id"`
	CheckSSL        int    `json:"check_ssl" gorm:"column:check_ssl"`
	CheckSSLAdvance int    `json:"check_ssl_advance" gorm:"column:check_ssl_advance"`
}

func GetAllActiveSSLList() []RActiveSSLItem {
	var result []RActiveSSLItem
	sqlStr := `
	SELECT
	ssl_config.id,ssl_config.web_id,web_list.group_id,monitor_policy.check_ssl,monitor_policy.check_ssl_advance
FROM
	ssl_config
LEFT JOIN web_list on web_list.id=ssl_config.web_id
LEFT JOIN monitor_policy ON monitor_policy.id = web_list.policy_id
WHERE
	ssl_config.deleted_at IS NULL`
	DB.Raw(sqlStr).Scan(&result)
	return result
}
