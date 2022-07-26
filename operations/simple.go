package operations

func Add(nums ...int) int {
	var result int

	for _, val := range nums {
		result += val
	}
	return result
}

func Subtract(x, y int) int {
	return x - y
}
