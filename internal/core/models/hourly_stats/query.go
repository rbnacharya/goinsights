package hourlystats

import (
	"time"
)

func (repo HourlyStatsRepository) FindHourlyStatsByCustomerID(customerID int, startDate int64, endDate int64) (*HourlyStats, *error) {
	// SQL query to find hourly stats by customer ID
	query := "SELECT sum(request_count), sum(invalid_count), customer_id FROM hourly_stats WHERE customer_id = $1 AND time >= $2 AND time <= $3 group by customer_id"

	startDateUnix := time.Unix(startDate, 0)
	endDateUnix := time.Unix(endDate, 0)

	rows, err := repo.Global.DB.Query(query, customerID, startDateUnix, endDateUnix)

	if err != nil {
		println(err.Error())
		return nil, &err
	}

	defer rows.Close()

	var hourlyStats []HourlyStats

	for rows.Next() {
		hourlyStat := HourlyStats{}
		err := rows.Scan(&hourlyStat.RequestCount, &hourlyStat.InvalidCount, &hourlyStat.CustomerID)
		if err != nil {
			return nil, &err
		}
		hourlyStats = append(hourlyStats, hourlyStat)
	}
	if len(hourlyStats) == 0 {
		return nil, nil
	} else {
		return &hourlyStats[0], nil
	}
}

func (repo HourlyStatsRepository) AddEntry(customerId int, timestamp int64, invalidRequest bool) error {
	// SQL query to insert entry
	db := repo.Global.DB

	// Convert timestamp to hourly timestamp

	startOfLastHourTimestamp := (int64(timestamp/3600)*3600 - 3600)

	convertedTimestamp := time.Unix(startOfLastHourTimestamp, 0)

	var query string

	if invalidRequest {
		query = "INSERT INTO hourly_stats (customer_id, time, request_count, invalid_count) VALUES ($1, $2, 1, 1) ON CONFLICT (customer_id, time) DO UPDATE SET request_count = hourly_stats.request_count + 1, invalid_count = hourly_stats.invalid_count + 1"
	} else {
		query = "INSERT INTO hourly_stats (customer_id, time, request_count, invalid_count) VALUES ($1, $2, 1, 0) ON CONFLICT (customer_id, time) DO UPDATE SET request_count = hourly_stats.request_count + 1"
	}

	// Convert timestamp to hourly timestamp
	_, err := db.Exec(query, customerId, convertedTimestamp)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
