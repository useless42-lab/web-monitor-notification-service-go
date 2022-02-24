package jobs

import (
	"WebMonitor/tools"
	"WebMonitor/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func PushNotification(data tools.RDeviceItemJob, eventJob *tgbotapi.BotAPI) {
	if data.NotificationItem.NotificationType == 1 {
		// 邮件
		data.NotificationItem.NotificationTarget = data.NotificationItem.Email
		utils.SendMail(data)
	} else if data.NotificationItem.NotificationType == 4 {
		// telegram
		data.NotificationItem.NotificationTarget = data.NotificationItem.Telegram
		utils.PushTelegram(eventJob, data)
	} else if data.NotificationItem.NotificationType == 5 {
		// bark
		data.NotificationItem.NotificationTarget = data.NotificationItem.Bark
		utils.PushBark(data)
	} else if data.NotificationItem.NotificationType == 6 {
		// server酱
		data.NotificationItem.NotificationTarget = data.NotificationItem.ServerChan
		utils.PushServerChan(data)
	}
}
