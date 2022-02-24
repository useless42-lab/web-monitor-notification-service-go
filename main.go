package main

import (
	"WebMonitor/jobs"
	"WebMonitor/utils"
	"fmt"
	"log"
	"os"

	"github.com/bamzi/jobrunner"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	go utils.InitTelegramBot(bot)

	jobrunner.Start()
	jobrunner.Schedule("@every 15s", jobs.CheckWebJob{Bot: bot})
	jobrunner.Schedule("@every 15s", jobs.CheckServerJob{Bot: bot})
	jobrunner.Schedule("@every 15s", jobs.CheckApiJob{Bot: bot})
	jobrunner.Schedule("@every 15s", jobs.CheckTcpJob{Bot: bot})
	jobrunner.Schedule("@every 15s", jobs.CheckDnsJob{Bot: bot})
	jobrunner.Schedule("@every 15s", jobs.CheckHeartbeatJob{Bot: bot})
	jobrunner.Schedule("@every 15s", jobs.CheckSteamGameServerJob{Bot: bot})
	jobrunner.Schedule("@every 15s", jobs.CheckMinecraftServerJob{Bot: bot})
	jobrunner.Schedule("@daily", jobs.CheckSSLJob{Bot: bot})
	jobrunner.Schedule("@daily", jobs.CheckWHOISJob{Bot: bot})

	var str string
	fmt.Scan(&str)
}
