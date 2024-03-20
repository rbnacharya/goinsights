package validation

import "time"

func (ctrl Validate) ValidateTimestamp(timestamp int64) bool {
	currentTimestamp := time.Now().Unix()
	// Check if timestamp is valid
	// Timestamp should not be in future
	// Timestamp should not be older than 1 year
	valid := timestamp < currentTimestamp

	if valid {
		valid = timestamp > (currentTimestamp - (365 * 24 * 60 * 60))
	}

	return valid
}
