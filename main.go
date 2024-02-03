package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	
	result := add(3, 7)
	fmt.Println("Sum:", result)
	
	// Check for division by zero
	if num2 == 0 {
	fmt.Println("Error: Division by zero is not allowed.") return
	}
	// Call the function and store the return values
	quotient, remainder := divide(num1, num2)
	// Print the results
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder) //End of Sindhuja Peravali - 500228575 code

}
func add(a, b int) int {
	return a + b
}
func revString(str string) string {
	var theArray []string
	theArray = strings.Split(str, "")
	var strOutput string
	for i := len(str) - 1; i >= 0; i-- {
		strOutput += theArray[i]
	}
	return strOutput
}
func divide(dividend, divisor int) (quotient, remainder int) {
	// Calculate the quotient 
	quotient = dividend / divisor
	// Calculate the remainder 
	remainder = dividend % divisor 
	return
}
