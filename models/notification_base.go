package models

func UpdateNotificationTelegram(userId int64, telegramId int64) {
	sqlStr := `update notification_base set telegram=@telegram where user_id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{
		"telegram": telegramId,
		"userId":   userId,
	})
}
