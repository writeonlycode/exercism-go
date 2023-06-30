package lasagna

// PreparationTime returns the amount of time to prepare the lasagna.
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	return len(layers) * time
}

// Quantitites returns the amount of noodles and sauce needed to prepare the
// lasagna.
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50

		case "sauce":
			sauce += 0.2
		}
	}

	return
}

// AddSecretIngredient replaces the last element in our recipe with the last
// element of our friend's recipe. The function mutates our recipe.
func AddSecretIngredient(friendList []string, myList []string) {
	myList = append(myList[:len(myList)-1], friendList[len(friendList)-1])
}

// ScaleRecipe returns the amounts of ingredients necessary to make the given
// number of portions of the recipe. It assumes that the original quantities
// are for 2 portions.
func ScaleRecipe(original []float64, portions int) []float64 {
	var scaled []float64

	for i := 0; i < len(original); i++ {
		scaled = append(scaled, (original[i]/2)*float64(portions))
	}

	return scaled
}
