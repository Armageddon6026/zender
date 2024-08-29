package common

import "time"

const (
	SYSTEM_CONFIG_PATH string        = "config/server.json"
	API_BASE_PATH      string        = "/api/v1"
	API_SSE_PATH       string        = "/events"
	JWT_HEADER         string        = "x-access-token"
	JWT_ISSUER         string        = "Armageddon"
	JWT_EXPIRE_TIME    time.Duration = 3 * 24 * time.Hour
)
