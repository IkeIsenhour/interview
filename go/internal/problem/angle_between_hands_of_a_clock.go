package problem

func AngleClock(hour int, minutes int) float64 {
	minuteTicks := float64(minutes)
	hourTicks := (float64(hour) * 5) + ((minuteTicks / 60) * 5)

	var diff float64
	if minuteTicks > hourTicks {
		diff = minuteTicks - hourTicks
	} else {
		diff = hourTicks - minuteTicks
	}

	angleOne := diff * 6
	angleTwo := 360 - angleOne

	return min(angleOne, angleTwo)
}
