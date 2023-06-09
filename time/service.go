package time

// FullGameTimeInMilliseconds is the full game time in milliseconds (10 minutes)
const FullGameTimeInMilliseconds = 600000

// ConvertSecondsToMilliseconds converts seconds to milliseconds
func ConvertSecondsToMilliseconds(timeInSeconds float64) (timeInMilliseconds float64) {
	return timeInSeconds * 1000
}

func GetElapsedTimeFromRemainingTime(remainingTimeInSeconds int) (elapsedTimeInMilliseconds float64) {
	return FullGameTimeInMilliseconds - ConvertSecondsToMilliseconds(float64(remainingTimeInSeconds))
}

// GetNewCooldown returns the new cooldown after applying a cooldown reduction
func GetNewCooldown(originalCooldown float64, reductionPercent float64) (newCooldown float64) {
	newCooldown = originalCooldown * (1 - reductionPercent)
	return
}
