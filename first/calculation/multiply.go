package calculation

func Multiply(num int, num2 int) int {
	if checkNum(num) {
		return num * num2
	}
	return 0
}

func checkNum(num int) bool {
	return num > 0
}
