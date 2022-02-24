package models

import "time"

func GetMonitorMinecraftServerSuccessPercent(times int, id string) SuccessPercentItem {
	var result SuccessPercentItem
	sqlStr := `
	SELECT
	CONCAT( CEILING( sum( a.check_success ) / @times ), "", "" ) AS percent 
FROM
	( SELECT check_success FROM minecraft_server_log WHERE minecraft_server_id = @id ORDER BY id DESC LIMIT @times ) AS a
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"times": times,
		"id":    id,
	}).Scan(&result)
	return result
}

type MinecraftServerLogItem struct {
	DefaultModel
	MinecraftServerId string `json:"minecraft_server_id" gorm:"column:minecraft_server_id"`
	Description       string `json:"description" gorm:"column:description"`
	Favicon           string `json:"favicon" gorm:"column:favicon"`
	PlayersMax        int    `json:"players_max" gorm:"column:players_max"`
	PlayersOnline     int    `json:"players_online" gorm:"column:players_online"`
	VersionName       string `json:"version_name" gorm:"column:version_name"`
	VersionProtocol   int    `json:"version_protocol" gorm:"column:version_protocol"`
	Region            int    `json:"region" gorm:"column:region"`
}

type RMinecraftServerLogItem struct {
	Id                string    `json:"id" gorm:"column:id"`
	MinecraftServerId string    `json:"minecraft_server_id" gorm:"column:minecraft_server_id"`
	Description       string    `json:"description" gorm:"column:description"`
	Favicon           string    `json:"favicon" gorm:"column:favicon"`
	PlayersMax        int       `json:"players_max" gorm:"column:players_max"`
	PlayersOnline     int       `json:"players_online" gorm:"column:players_online"`
	VersionName       string    `json:"version_name" gorm:"column:version_name"`
	VersionProtocol   int       `json:"version_protocol" gorm:"column:version_protocol"`
	Region            int       `json:"region" gorm:"column:region"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
}

func AddMinecraftServerLog(minecraftServerId string, description string, favicon string, playersMax int, playersOnline int, versionName string, versionProtocol int, region int) {
	data := MinecraftServerLogItem{
		MinecraftServerId: minecraftServerId,
		Description:       description,
		Favicon:           favicon,
		PlayersMax:        playersMax,
		PlayersOnline:     playersOnline,
		VersionName:       versionName,
		VersionProtocol:   versionProtocol,
		Region:            region,
	}
	DB.Table("minecraft_server_log").Create(&data)
}

func GetLatestMinecraftServerLog(minecraftServerId string) RMinecraftServerLogItem {
	var result RMinecraftServerLogItem
	sqlStr := `select * from minecraft_server_log where minecraft_server_id=@minecraftServerId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"minecraftServerId": minecraftServerId,
	}).Scan(&result)
	return result
}
