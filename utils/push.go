package utils

import (
	"WebMonitor/tools"
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type RServerChan struct {
	Text string
	Desp string
}

func PushServerChan(data tools.RDeviceItemJob) {
	inputData := url.Values{
		"text": {data.Name},
		"desp": {"出现了些许问题！"},
	}
	clt := http.Client{}
	clt.PostForm(data.NotificationItem.NotificationTarget, inputData)
}

type RPushPlusItem struct {
	Icon        string    `json:"icon" gorm:"column:icon"`
	Title       string    `json:"title" gorm:"column:title"`
	Link        string    `json:"link" gorm"column:link"`
	Description string    `json:"description" gorm:"column:description"`
	TargetTime  time.Time `json:"target_time" gorm:"column:target_time"`
}

func PushPlus(data tools.RDeviceItemJob) {
	newData := map[string]interface{}{
		"name": data.Name,
		"desp": "出现了些许问题！",
	}

	jsonStr, _ := json.Marshal(newData)
	notificationTarget, _ := regexp.Compile("http://")
	notificationTarget2 := notificationTarget.ReplaceAllString(data.NotificationItem.NotificationTarget, "https://")
	targetUrl := notificationTarget2 + "&title=" + data.Name + "&content=" + url.QueryEscape(string(jsonStr)) + "&template=json"

	resp, err := http.Get(targetUrl)
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()

}

func PushWebhook(data tools.RDeviceItemJob) {
	inputData := url.Values{
		"name": {data.Name},
		"desp": {"出现了些许问题！"},
	}
	clt := http.Client{}
	clt.PostForm(data.NotificationItem.NotificationTarget, inputData)
}

func PushBark(data tools.RDeviceItemJob) {
	clt := http.Client{}
	clt.Get(data.NotificationItem.NotificationTarget + "/" + data.Name + "/出现了些许问题！")
}
