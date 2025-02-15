package circle

import "math"

func GetCircleCircumference(radius float64) float64 {
	return 2 * math.Phi * radius
}

func GetCircleArea(radius float64) float64 {
	return math.Phi * math.Pow(radius, 2)
}
