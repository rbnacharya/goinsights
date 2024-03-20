package ipuablacklist

func (repo IPUABlacklistRepository) FindIPBlacklistByIP(ip string) *IPBlacklist {
	// SQL query to find IP blacklist by IP
	query := `SELECT ip FROM ip_blacklist WHERE ip = $1`
	db := repo.Global.DB

	row := db.QueryRow(query, ip)

	ipBlacklist := IPBlacklist{}
	err := row.Scan(&ipBlacklist.IP)
	if err != nil {
		println(err.Error())
		return nil
	}

	return &ipBlacklist
}
