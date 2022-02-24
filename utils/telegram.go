package utils

import (
	"WebMonitor/models"
	"WebMonitor/tools"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	_ "github.com/joho/godotenv/autoload"
)

func InitTelegramBot(bot *tgbotapi.BotAPI) {
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		chatId := update.Message.Chat.ID
		msgArr := strings.Fields(msg.Text)
		if len(msgArr) > 1 {
			eventId, _ := strconv.ParseInt(msgArr[1], 10, 64)
			// models.UpdateNotificationTarget(eventId, strconv.FormatInt(chatId, 10))
			models.UpdateNotificationTelegram(eventId, chatId)
			msg.Text = "绑定成功"

		} else {
			msg.Text = "https://www.pingsilo.com/"
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func PushTelegram(bot *tgbotapi.BotAPI, data tools.RDeviceItemJob) {
	chatId, _ := strconv.ParseInt(data.NotificationItem.NotificationTarget, 10, 64)

	messageConfig := tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{ChatID: chatId},
		Text:                  data.Name + " " + data.Path,
		ParseMode:             "Markdown",
		DisableWebPagePreview: false,
	}
	bot.Send(messageConfig)
}
