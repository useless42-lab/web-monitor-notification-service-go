package models

import "time"

func GetMonitorSteamGameServerSuccessPercent(times int, id string) SuccessPercentItem {
	var result SuccessPercentItem
	sqlStr := `
	SELECT
	CONCAT( CEILING( sum( a.check_success ) / @times ), "", "" ) AS percent 
FROM
	( SELECT check_success FROM steam_game_server_log WHERE steam_game_server_id = @id ORDER BY id DESC LIMIT @times ) AS a
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"times": times,
		"id":    id,
	}).Scan(&result)
	return result
}

type SteamGameServerLogItem struct {
	DefaultModel
	SteamGameServerId string `json:"steam_game_server_id" gorm:"column:steam_game_server_id"`
	Name              string `json:"name" gorm:"column:name`
	PlayersMax        int    `json:"players_max" gorm:"column:players_max"`
	PlayersOnline     int    `json:"players_online" gorm:"column:players_online"`
}

type RSteamGameServerLogItem struct {
	Id                string    `json:"id" gorm:"column:id"`
	SteamGameServerId string    `json:"steam_game_server_id" gorm:"column:steam_game_server_id"`
	Name              string    `json:"name" gorm:"column:name`
	PlayersMax        int       `json:"players_max" gorm:"column:players_max"`
	PlayersOnline     int       `json:"players_online" gorm:"column:players_online"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
}

func AddSteamGameServerLog(steamGameServerId string, name string, playersMax int, playersOnline int) {
	data := SteamGameServerLogItem{
		SteamGameServerId: steamGameServerId,
		Name:              name,
		PlayersMax:        playersMax,
		PlayersOnline:     playersOnline,
	}
	DB.Table("steam_game_server_log").Create(&data)
}

func GetLatestSteamGameServerLog(steamGameServerId string) RSteamGameServerLogItem {
	var result RSteamGameServerLogItem
	sqlStr := `select * from steam_game_server_log where steam_game_server_id=@steamGameServerId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"steamGameServerId": steamGameServerId,
	}).Scan(&result)
	return result
}
