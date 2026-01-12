package constants

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

// loaded from .env using autoload
var (
	LAYOUT_PATH			 	= os.Getenv("LAYOUT_PATH")
	NB_AGENTS 				= envInt("NB_AGENTS", 20)

	AGENT_SPAWN_INTERVAL 	= envInt("AGENT_SPAWN_INTERVAL", 150)
	INCREMENTAL_VELOCITY	= envFloat("INCREMENTAL_VELOCITY", 0.005)
	INCREMENTAL_DIRECTION 	= envFloat("INCREMENTAL_DIRECTION", 0.05)
	MIN_VELOCITY 			= envFloat("MIN_VELOCITY", 0.1)
	MAX_VELOCITY 			= envFloat("MAX_VELOCITY", 1.0)

	TIC_DURATION 			= envInt("TIC_DURATION", 30)
)

func envInt(key string, def int) int {
	if v, _ := strconv.Atoi(os.Getenv(key)); v >= 0 {
		return v
	}
	return def
}

func envFloat(key string, def float64) float64 {
	if v, _ := strconv.ParseFloat(os.Getenv(key), 64); v > 0 {
		return v
	}
	return def
}
