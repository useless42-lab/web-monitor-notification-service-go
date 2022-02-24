package models

import "time"

type WHOISLItem struct {
	DomainExpirationDate time.Time `json:"domain_expiration_date" gorm:"column:domain_expiration_date"`
}

func GetWHOISConfig(id int64) WHOISLItem {
	var result WHOISLItem
	sqlStr := `select domain_expiration_date from domain_whois where id=@id and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"id": id,
	}).Scan(&result)
	return result
}

type RActiveWHOISItem struct {
	Id                int64  `json:"id" gorm:"column:id"`
	WebId             int64  `json:"web_id" gorm:"column:web_id"`
	Name              string `json:"name" gorm:"column:name"`
	Path              string `json:"path" gorm:"column:path"`
	GroupId           int64  `json:"group_id" gorm:"column:group_id"`
	CheckWHOIS        int    `json:"check_whois" gorm:"column:check_whois"`
	CheckWHOISAdvance int    `json:"check_whois_advance" gorm:"column:check_whois_advance"`
}

func GetAllActiveWHOISList() []RActiveWHOISItem {
	var result []RActiveWHOISItem
	sqlStr := `
	SELECT
	domain_whois.id,domain_whois.web_id,web_list.group_id,monitor_policy.check_whois,monitor_policy.check_whois_advance
FROM
	domain_whois
LEFT JOIN web_list on web_list.id=domain_whois.web_id
LEFT JOIN monitor_policy ON monitor_policy.id = web_list.policy_id
WHERE
	domain_whois.deleted_at IS NULL`
	DB.Raw(sqlStr).Scan(&result)
	return result
}
