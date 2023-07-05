package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	myMap["apple"] = 1
	myMap["bannana"] = 2
	myMap["orange"] = 3

	appleValue := myMap["apple"]
	bannanaValue := myMap["bannana"]
	fmt.Println("value of apple:", appleValue)
	fmt.Println("value of bannana:", bannanaValue)

	myMap["apple"] = 5

	fmt.Println("updated value of apple:", myMap["apple"])

	delete(myMap, "orange")
	fmt.Println("after deleting orange:", myMap)
	value, exits := myMap["banana"]
	if exits {
		fmt.Println("value of bannana:", value)
	} else {
		fmt.Println("bannaana not found in the map")

	}
	for key, value := range myMap {
		fmt.Println("key:", key, "Value:", value)
	}
	fmt.Println("length of the map:", len(myMap))
}
