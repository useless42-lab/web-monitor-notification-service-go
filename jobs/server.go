package jobs

import (
	"WebMonitor/models"
	"WebMonitor/tools"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CheckServerJob struct {
	Bot *tgbotapi.BotAPI
}

func (checkServerJob CheckServerJob) Run() {
	var deviceType int = 2
	serverList := models.GetAllActiveServerList()
	for _, item := range serverList {
		// 1 ping 2 服务器状态
		successPercent := models.GetMonitorServerSuccessPercent(item.FailedWaitTimes, item.Id)
		if successPercent.Percent == 0 {
			// 监控失败 发送通知
			deviceId, _ := strconv.ParseInt(item.Id, 10, 64)
			notificationLog := models.GetLatestNotificationLog(deviceType, deviceId)
			targetNotification := models.GetTargetNotificationList(item.GroupId, deviceType)
			if time.Now().After(notificationLog.CreatedAt.Add(+time.Hour * time.Duration(24))) {
				for _, item2 := range targetNotification {
					data := tools.RDeviceItemJob{
						Name: item.Name,
						Path: item.Path,
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

					go PushNotification(data, checkServerJob.Bot)
					models.AddNotificationLog(notificationId, deviceType, deviceId)
				}
			}
		}
	}
}
