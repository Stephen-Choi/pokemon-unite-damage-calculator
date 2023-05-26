package time

// ConvertSecondsToMilliseconds converts seconds to milliseconds
func ConvertSecondsToMilliseconds(timeInSeconds float64) (timeInMilliseconds float64) {
	return timeInSeconds * 1000
}

// GetNewCooldown returns the new cooldown after applying a cooldown reduction
func GetNewCooldown(originalCooldown float64, reductionPercent float64) (newCooldown float64) {
	newCooldown = originalCooldown * (1 - reductionPercent)
	return
}
