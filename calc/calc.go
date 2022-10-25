package calc

func ARB(frontPercent float64) (float64, float64) {
	front := frontPercent / 100
	rear := 1 - front
	frontVal := (65-1)*front + 1 + (12.0 / 65.0)
	rearVal := (65-1)*rear +  1 + (12.0 / 65.0)

	return frontVal, rearVal
}

func Springs(frontPercent, min, max float64) (float64, float64) {
	front := frontPercent / 100
	rear := 1 - front
	frontVal := (max - min) * front + min
	rearVal := (max-min) * rear + min

	return frontVal, rearVal
}

func Damping(frontPercent, min, max float64) (float64, float64, float64, float64) {
	front := frontPercent / 100
	rear := 1 - front
	reboundFrontVal := (max-min)*front + min
	reboundRearVal := (max-min)*rear + min
	bumpFrontVal := reboundFrontVal * 0.625
	bumpRearVal := reboundRearVal * 0.625

	return reboundFrontVal, reboundRearVal, bumpFrontVal, bumpRearVal
}
