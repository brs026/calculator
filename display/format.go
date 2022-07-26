package display

import (
	"fmt"

	"github.com/brs026/calculator/operations"
)

// Takes a symbol like +, *, -, ...,
// An initial value from which to begin the calculation,
// a function to apply; in case of addition x + y the function to apply to 2 values
//func printOperation(symbol string, initialValue int, fn func(x, y int) int, operands ...int) {
//	//result += initialValue
//	fmt.Println("The opetion is:", symbol)/
//	fmt.Println("The init value:", initialValue)
//	fmt.Println("Operands:", operands)
//}

//func PrintAddition(fn func(x, y int) int, numbers ...int) {
func PrintAddition(numbers ...int) int {
	//printOperation("+", 0, fn, numbers...)
	result := operations.Add(numbers...)
	fmt.Println("Sum of:", numbers, "= ", result)
	return result
}

//func PrintMultiply(fn func(x, y int) int, numbers ...int) {
//	printOperation("*", 1, fn, numbers...)
//}

func printDemarcation() {
	demarcation := ""

	for index := 0; index < 15; index++ {
		demarcation += "-"
	}
	fmt.Println(demarcation)
}

// format operands into the same length, so they are always aligned
// Example:
//	 3
// +10
func makeSameLength(x, y string) string {
	difference := len(x) - len(y)

	if difference > 0 {
		return x
	}

	difference *= -1

	result := ""

	for index := difference; index > 0; index-- {
		result += " "
	}
	return result + x
}
