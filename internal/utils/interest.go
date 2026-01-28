package utils

// The lower, the better
func NormMin(value, min, max float64) float64 {
	if value <= min {
		return 1
	}
	if value >= max {
		return 0
	}
	return (max - value) / (max - min)
}

// The higher, the better
func NormMax(value, min, max float64) float64 {
	if value <= min {
		return 0
	}
	if value >= max {
		return 1
	}
	return (value - min) / (max - min)
}

func BoolBonus(v bool, bonus float64) float64 {
	if v {
		return bonus
	}
	return 0
}
