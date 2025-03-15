package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, layer := range layers {
		if layer == "noodles" {
			noodles++
		} else if layer == "sauce" {
			sauce++
		}
	}
	noodles *= 50
	sauce *= 0.2
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ing1, ing2 []string) {
	var secretIngredient = ing1[len(ing1)-1]
	ing2[len(ing2)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(baseAmounts []float64, number int) []float64 {
	portions := make([]float64, len(baseAmounts))
	for i, v := range baseAmounts {
		portions[i] = v * float64(number) / 2.0
	}
	return portions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
