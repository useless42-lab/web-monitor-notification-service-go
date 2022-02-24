package models

type RNotificationItemStruct struct {
	Id               string `json:"id" gorm:"column:id"`
	NotificationType int    `json:"notification_type" gorm:"column:notification_type"`
	Email            string `json:"email" gorm:"column:email"`
	Phone            string `json:"phone" gorm:"column:phone"`
	SMS              string `json:"sms" gorm:"column:sms"`
	Telegram         string `json:"telegram" gorm:"column:telegram"`
	Bark             string `json:"bark" gorm:"gorm:bark"`
	ServerChan       string `json:"server_chan" gorm:"column:server_chan"`
}

func GetTargetNotificationList(groupId int64, deviceType int) []RNotificationItemStruct {
	var result []RNotificationItemStruct
	sqlStr := `
	SELECT
	nl.id,nl.notification_type,nb.email,nb.phone,nb.sms,nb.telegram,nb.bark,nb.server_chan 
FROM
	notification_list AS nl
	LEFT JOIN notification_base AS nb ON nl.user_id = nb.user_id 
WHERE
	nl.group_id = @groupId 
	AND nl.device_type = @deviceType 
	AND nl.deleted_at IS NULL
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"groupId":    groupId,
		"deviceType": deviceType,
	}).Scan(&result)
	return result
}
