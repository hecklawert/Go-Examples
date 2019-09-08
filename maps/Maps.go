package main

import "fmt"

func main() {
	// Define our map with the population of some cities
	statePopulation := map[string]int{
		"Madrid":    2313345,
		"Bilbao":    1432659,
		"Barcelona": 2184254,
		"Logrono":   145931,
	}
	pop, ok := statePopulation["Logrono"]
	fmt.Println(pop, ok)
	pop, ok = statePopulation["Valencia"]
	fmt.Println(pop, ok)
}
