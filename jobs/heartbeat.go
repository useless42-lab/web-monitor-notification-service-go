package jobs

import (
	"WebMonitor/models"
	"WebMonitor/tools"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CheckHeartbeatJob struct {
	Bot *tgbotapi.BotAPI
}

func (checkHeartbeatJob CheckHeartbeatJob) Run() {
	var deviceType int = 6
	heartbeatList := models.GetAllActiveHeartbeatList()
	for _, item := range heartbeatList {
		successPercent := models.GetMonitorHeartbeatSuccessPercent(item.FailedWaitTimes, item.Id)
		if successPercent.Percent == 0 {
			// 监控失败 发送通知
			deviceId, _ := strconv.ParseInt(item.Id, 10, 64)
			notificationLog := models.GetLatestNotificationLog(deviceType, deviceId)
			if time.Now().After(notificationLog.CreatedAt.Add(+time.Hour * time.Duration(24))) {
				targetNotification := models.GetTargetNotificationList(item.GroupId, deviceType)
				for _, item2 := range targetNotification {
					data := tools.RDeviceItemJob{
						Name: item.Name,
						Path: item.Token,
						NotificationItem: tools.RNotificationItemStruct{
							NotificationType: item2.NotificationType,
							Email:            item2.Email,
							Phone:            item2.Phone,
							SMS:              item2.SMS,
							Telegram:         item2.Telegram,
							Bark:             item2.Bark,
							ServerChan:       item2.ServerChan,
						},
					}
					notificationId, _ := strconv.ParseInt(item2.Id, 10, 64)

					go PushNotification(data, checkHeartbeatJob.Bot)
					models.AddNotificationLog(notificationId, deviceType, deviceId)
				}
			}
		}
	}
}
