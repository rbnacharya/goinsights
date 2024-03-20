package ipuablacklist

import "github.com/rbnacharya/goinsights/internal/core"

type IPBlacklist struct {
	IP string `json:"ip"`
}

type UABlacklist struct {
	UA string `json:"ua"`
}

type IPUABlacklistRepository struct {
	Global *core.Global
}
