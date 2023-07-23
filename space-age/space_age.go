package space

type Planet string

var planets = map[Planet]float64{
	Planet("Mercury"): 0.2408467,
	Planet("Venus"):   0.61519726,
	Planet("Earth"):   1.0,
	Planet("Mars"):    1.8808158,
	Planet("Jupiter"): 11.862615,
	Planet("Saturn"):  29.447498,
	Planet("Uranus"):  84.016846,
	Planet("Neptune"): 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	value, exists := planets[planet]

	if !exists {
		return -1.0
	}

	return seconds / (31557600 * value)
}
