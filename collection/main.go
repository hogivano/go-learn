package main

import "fmt"

func main() {
	fmt.Println(names())
	fmt.Println(ages())
	fmt.Println(slice())

	maps()
}

func names() [5]string {
	var names [5]string

	names[0] = "Hend"
	names[1] = "Dri"
	names[2] = "Ogi"
	names[3] = "Vano"
	names[4] = "Hen"

	return names
}

func ages() [9]int {
	ages := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	return ages
}

func slice() []string {
	//slice is for undefined length of array
	brands := []string{"Logitect", "Keycron"}

	//push new value
	brands = append(brands, "Hend")

	return brands
}

func maps() {
	var myMaps map[int]string
	myMaps = map[int]string{}
	myMaps[100] = "oke"

	for index, item := range myMaps {
		fmt.Println(index, item)
	}
}
