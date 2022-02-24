package models

import "time"

type TcpLogItem struct {
	DefaultModel
	TcpId        string `json:"tcp_id" gorm:"column:tcp_id"`
	Elapsed      int    `json:"elapsed" gorm:"column:elapsed"`
	CheckSuccess int    `json:"check_success" gorm:"column:check_success"`
	Region       int    `json:"region" gorm:"column:region"`
}

type RTcpLogItem struct {
	TcpId        string    `json:"tcp_id" gorm:"column:tcp_id"`
	CheckSuccess int       `json:"check_success" gorm:"column:check_success"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
}

func AddTcpLog(tcpId string, elapsed int, checkSuccess int, region int) {
	data := TcpLogItem{
		TcpId:        tcpId,
		Elapsed:      elapsed,
		CheckSuccess: checkSuccess,
		Region:       region,
	}
	DB.Table("tcp_log").Create(&data)
}

func GetLatestTcpLog(tcpId string) RTcpLogItem {
	var result RTcpLogItem
	sqlStr := `select * from tcp_log where tcp_id=@tcpId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"tcpId": tcpId,
	}).Scan(&result)
	return result
}

func GetMonitorTcpSuccessPercent(times int, id string) SuccessPercentItem {
	var result SuccessPercentItem
	sqlStr := `
	SELECT
	CONCAT( CEILING( sum( a.check_success ) / @times ), "", "" ) AS percent 
FROM
	( SELECT check_success FROM tcp_log WHERE tcp_id = @id ORDER BY id DESC LIMIT @times ) AS a
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"times": times,
		"id":    id,
	}).Scan(&result)
	return result
}
