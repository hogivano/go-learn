package main

import "fmt"

func main() {
	var name string
	name = "Hendriyan Ogivano"
	age := 20

	fmt.Println(age, name)
	fmt.Println(switchCon(10))

	quiz()
}

func condition() string {
	age := 23

	if age <= 20 {
		return "You're young"
	} else if age >= 30 {
		return "You're old"
	}
	return "You're a man"
}

func switchCon(num int) string {
	var ages string
	switch num {
	case 10:
		ages = "Kids"
	case 18:
		ages = "Teen"
	case 20:
		ages = "Man"
	default:
		ages = "Undinfined"
	}

	return ages
}

func loops() {
	for i := 0; i < 20; i++ {
		fmt.Println("index : ", i+1)
	}

	//like while
	i := 1
	for i < 10 {
		fmt.Println("index : ", i)
		i++
	}

	name := "Hendriyan"

	for _, letter := range name {
		fmt.Println("letter : ", string(letter))
	}
}

func quiz() {
	title := "Golang is the best language"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index : ", index, " letter : ", string(letter))
		}
	}

	for _, letter := range title {
		switch string(letter) {
		case "a", "i", "u", "e", "o":
			fmt.Println("Char : ", string(letter))
		}
	}
}
