package models

type RSteamGameServerItem struct {
	Id              string    `json:"id" gorm:"column:id"`
	Name            string    `json:"name" gorm:"column:name"`
	Path            string    `json:"path" gorm:"column:path"`
	GroupId         int64     `json:"group_id" gorm:"column:group_id"`
	Frequency       int       `json:"frequency" gorm:"column:frequency"`
	FailedWaitTimes int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt       LocalTime `json:"created_at" gorm:"column:created_at"`
	Status          int       `json:"status" gorm:"column:status"`
	SteamApiKey     string    `json:"steam_api_key" gorm:"steam_api_key"`
}

func GetAllActiveSteamGameServerList() []RSteamGameServerItem {
	var result []RSteamGameServerItem
	sqlStr := `
	SELECT
	d.id,
	d.name,
	d.path,
	d.group_id,
	d.status,
	d.created_at,
	mp.frequency,
	mp.failed_wait_times,
	u.steam_api_key
FROM
	steam_game_server_list as d
LEFT JOIN monitor_policy as mp ON mp.id = d.policy_id
INNER JOIN user as u on u.id=d.user_id
WHERE
	d.deleted_at IS NULL
AND d.status = 1`
	DB.Raw(sqlStr).Scan(&result)
	return result
}
