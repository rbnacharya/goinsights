package register

import ipuablacklist "github.com/rbnacharya/trafficinsights-go/internal/core/models/ip_ua_blacklist"

func (ctrl Register) ValidateUserAgent(userAgent string) bool {
	// Check if user agent is in blacklist
	uaIpRepo := ipuablacklist.IPUABlacklistRepository{
		Global: ctrl.Global,
	}

	exists := uaIpRepo.FindUABlacklistByUA(userAgent) != nil

	return exists
}
