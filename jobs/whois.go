package jobs

import (
	"WebMonitor/models"
	"WebMonitor/tools"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CheckWHOISJob struct {
	Bot *tgbotapi.BotAPI
}

func (checkWHOISJob CheckWHOISJob) Run() {
	var deviceType int = 10
	whoisList := models.GetAllActiveWHOISList()
	for _, item := range whoisList {
		if item.CheckWHOIS == 1 {
			deviceId := item.Id
			webWHOISConfig := models.GetWHOISConfig(deviceId)
			whoisTime := webWHOISConfig.DomainExpirationDate.Format("2006-01-02 15:04:05")
			whoisTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", whoisTime, time.Local)
			whoisTimeLocation = whoisTimeLocation.Add(-time.Second * time.Duration(item.CheckWHOISAdvance))
			if time.Now().After(whoisTimeLocation) {
				// WHOIS监控失败 发送通知
				notificationLog := models.GetLatestNotificationLog(deviceType, deviceId)
				if time.Now().After(notificationLog.CreatedAt.Add(+time.Hour * time.Duration(24))) {
					// 获取网站的通知列表
					targetNotification := models.GetTargetNotificationList(item.GroupId, 1)
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
						go PushNotification(data, checkWHOISJob.Bot)
						models.AddNotificationLog(notificationId, deviceType, deviceId)
					}
				}
			}
		}
	}
}
