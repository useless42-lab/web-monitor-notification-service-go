package models

import "time"

func GetMonitorHeartbeatSuccessPercent(times int, id string) SuccessPercentItem {
	var result SuccessPercentItem
	sqlStr := `
	SELECT
	CONCAT( CEILING( sum( a.check_success ) / @times ), "", "" ) AS percent 
FROM
	( SELECT check_success FROM heartbeat_log WHERE heartbeat_id = @id ORDER BY id DESC LIMIT @times ) AS a
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"times": times,
		"id":    id,
	}).Scan(&result)
	return result
}

type HeartbeatLogItem struct {
	DefaultModel
	HeartbeatId  string `json:"heartbeat_id" gorm:"column:heartbeat_id"`
	ResponseData string `json:"response_data" grom:"column:response_data"`
	CheckSuccess int    `json:"check_success" gorm:"check_success"`
}

type RHeartbeatLogItem struct {
	Id           string    `json:"id" gorm:"column:id"`
	HeartbeatId  string    `json:"heartbeat_id" gorm:"column:heartbeat_id"`
	ResponseData string    `json:"response_data" grom:"column:response_data"`
	CheckSuccess int       `json:"check_success" gorm:"check_success"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
}

func AddHeartbeatLog(heartbeatId string, responseData string, checkSuccess int) {
	data := HeartbeatLogItem{
		HeartbeatId:  heartbeatId,
		ResponseData: responseData,
		CheckSuccess: checkSuccess,
	}
	DB.Table("heartbeat_log").Create(&data)
}

func GetLatestHeartbeatLog(heartbeatId string) RHeartbeatLogItem {
	var result RHeartbeatLogItem
	sqlStr := `select * from heartbeat_log where heartbeat_id=@heartbeatId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"heartbeatId": heartbeatId,
	}).Scan(&result)
	return result
}
