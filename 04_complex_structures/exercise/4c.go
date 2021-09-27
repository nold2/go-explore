package main

import "fmt"

func main() {
	scores := make([]float64, 5, 10)
	scores[0] = 1.0
	scores[1] = 3.0
	scores[2] = 4.0
	scores[3] = 6.0
	scores[4] = 2.0

	average := func(list []float64) float64 {
		var sum float64
		for _, val := range list {
			sum += val
		}
		return sum / float64(len(list))
	}

	fmt.Println(average(scores))

	pets := map[string]string{
		"fido":     "dog",
		"callus":   "dragon",
		"gopher":   "sealion",
		"kurogiri": "manthis",
	}

	finder := func(name string) bool {
		_, ok := pets[name]
		return ok
	}

	fmt.Println(finder("FIDO"))
	fmt.Println(finder("fido"))

	groceries := []string{"flour", "egg", "milk", "cheese", "candy", "cereal"}
	cart := groceries[2:5]

	addToCart := func(list ...string) []string {
		result := cart
		for _, val := range list {
			result = append(result, val)
		}

		return result
	}

	fmt.Println(cart)
	fmt.Println(addToCart("mushroom", "eggplant", "apples", "bacon"))
}
