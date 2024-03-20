package register

// CREATE TABLE `hourly_stats` (
// 	`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
// 	`customer_id` int(11) unsigned NOT NULL,
// 	`time` timestamp NOT NULL,
// 	`request_count` bigint(20) unsigned NOT NULL DEFAULT '0',
// 	`invalid_count` bigint(20) unsigned NOT NULL DEFAULT '0',
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `unique_customer_time` (`customer_id`,`time`),
// 	KEY `customer_idx` (`customer_id`),
// 	CONSTRAINT `hourly_stats_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
//   );

func (ctrl Register) AddEntry(input RegisterClickInput, invalidRequest bool) error {
	// SQL query to insert entry
	db := ctrl.Global.DB

	var query string

	if invalidRequest {
		query = "INSERT INTO hourly_stats (customer_id, time, request_count, invalid_count) VALUES ($1, $2, 1, 1) ON CONFLICT (customer_id, time) DO UPDATE SET request_count = hourly_stats.request_count + 1, invalid_count = hourly_stats.invalid_count + 1"
	} else {
		query = "INSERT INTO hourly_stats (customer_id, time, request_count, invalid_count) VALUES ($1, $2, 1, 0) ON CONFLICT (customer_id, time) DO UPDATE SET request_count = hourly_stats.request_count + 1"
	}
	// Convert timestamp to hourly timestamp
	input.Timestamp = input.Timestamp / 3600 * 3600 // Convert to hourly timestamp
	_, err := db.Exec(query, input.CustomerID, input.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
