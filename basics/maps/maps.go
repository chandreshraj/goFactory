package main

import "fmt"

func main() {

	// studentage := make(map[string]int)

	// studentage["Araya"] = 10
	// studentage["chand"] = 22
	// studentage["Aji"] = 23
	// studentage["Ravi"] = 26
	// studentage["sharan"] = 27

	// fmt.Println(studentage["chand"])
	// fmt.Println(len(studentage))

	superHero := map[string]map[string]string{
		"Superman": {
			"RealName": "simon",
			"City":     "New York",
		},
		"Batmon": {
			"RealName": "Bruce",
			"City":     "New York",
		},
	}

	// fmt.Println(superHero)
	if temp, hero := superHero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
