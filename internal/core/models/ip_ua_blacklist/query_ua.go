package ipuablacklist

func (repo IPUABlacklistRepository) FindUABlacklistByUA(ua string) *UABlacklist {
	// SQL query to find IP blacklist by IP
	query := `SELECT ua FROM ua_blacklist WHERE ua = $1`
	db := repo.Global.DB

	row := db.QueryRow(query, ua)

	uaBlacklist := UABlacklist{}
	err := row.Scan(&uaBlacklist.UA)
	if err != nil {
		return nil
	}

	return &uaBlacklist
}
