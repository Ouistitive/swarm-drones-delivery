package constants

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

// loaded from .env using autoload
var (
	LAYOUT_PATH = os.Getenv("LAYOUT_PATH")
	ADDRESSES_PATH = os.Getenv("ADDRESSES_PATH")
	NB_AGENTS   = envInt("NB_AGENTS", 20)

	AGENT_SPAWN_INTERVAL  = envInt("AGENT_SPAWN_INTERVAL", 150)
	INCREMENTAL_VELOCITY  = envFloat("INCREMENTAL_VELOCITY", 0.005)
	INCREMENTAL_DIRECTION = envFloat("INCREMENTAL_DIRECTION", 0.05)
	MIN_VELOCITY          = envFloat("MIN_VELOCITY", 0.1)
	MAX_VELOCITY          = envFloat("MAX_VELOCITY", 1.0)

	TIC_DURATION = envInt("TIC_DURATION", 30)

	BATTERY_DISCHARGING_MOVE    = envFloat("BATTERY_DISCHARGING_MOVE", 0.05)
	BATTERY_DISCHARGING_PICK    = envFloat("BATTERY_DISCHARGING_PICK", 1.0)
	BATTERY_DISCHARGING_DELIVER = envFloat("BATTERY_DISCHARGING_DELIVER", 1.0)
	BATTERY_CHARGING_RATE       = envFloat("BATTERY_CHARGING_RATE", 0.3)
	BATTERY_EMERGENCY_RATIO     = envFloat("BATTERY_EMERGENCY_RATIO", 0.5) // Below x% of battery, will try to recharge

	DISTANCE_PACKAGE_SCORE_WEIGHT     = envFloat("DISTANCE_PACKAGE_SCORE_WEIGHT", 0.2)
	DISTANCE_DESTINATION_SCORE_WEIGHT     = envFloat("DISTANCE_DESTINATION_SCORE_WEIGHT", 0.1)
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
