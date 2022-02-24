package jobs

import (
	"WebMonitor/models"
	"WebMonitor/tools"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CheckSSLJob struct {
	Bot *tgbotapi.BotAPI
}

func (checkSSLJob CheckSSLJob) Run() {
	var deviceType int = 9
	sslList := models.GetAllActiveSSLList()
	for _, item := range sslList {
		if item.CheckSSL == 1 {
			deviceId := item.Id
			webSSLConfig := models.GetSslConfig(deviceId)
			sslTime := webSSLConfig.TEndTime.Format("2006-01-02 15:04:05")
			sslTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", sslTime, time.Local)
			sslTimeLocation = sslTimeLocation.Add(-time.Second * time.Duration(item.CheckSSLAdvance))
			if time.Now().After(sslTimeLocation) {
				// SSL监控失败 发送通知
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
						go PushNotification(data, checkSSLJob.Bot)
						models.AddNotificationLog(notificationId, deviceType, deviceId)
					}
				}
			}
		}
	}
}
