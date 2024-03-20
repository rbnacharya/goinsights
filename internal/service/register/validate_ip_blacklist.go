package register

import ipuablacklist "github.com/rbnacharya/trafficinsights-go/internal/core/models/ip_ua_blacklist"

func (ctrl Register) ValidateIPBlacklist(ipAddress string) bool {
	// Check if IP address is in blacklist

	uaIpRepo := ipuablacklist.IPUABlacklistRepository{
		Global: ctrl.Global,
	}

	exists := uaIpRepo.FindIPBlacklistByIP(ipAddress) != nil
	return exists
}
