package circle

import "math"

func GetCircumference(radius float64) float64 {
	return 2 * math.Phi * radius
}

func GetArea(radius float64) float64 {
	return math.Phi * math.Pow(radius, 2)
}
