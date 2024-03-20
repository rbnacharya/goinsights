package validation

import ipuablacklist "github.com/rbnacharya/goinsights/internal/core/models/ip_ua_blacklist"

func (ctrl Validate) ValidateIPBlacklist(ipAddress string) bool {
	// Check if IP address is in blacklist

	uaIpRepo := ipuablacklist.IPUABlacklistRepository{
		Global: ctrl.Global,
	}

	exists := uaIpRepo.FindIPBlacklistByIP(ipAddress) != nil
	return exists
}
