package main

import "fmt"

func printReports(intro, body, outro string) {
	// Call printCostReport for the intro with an anonymous function to calculate its cost
	printCostReport(func(message string) int {
		return 2 * len(message)
	}, intro)

	// Call printCostReport for the body with an anonymous function to calculate its cost
	printCostReport(func(message string) int {
		return 3 * len(message)
	}, body)

	// Call printCostReport for the outro with an anonymous function to calculate its cost
	printCostReport(func(message string) int {
		return 4 * len(message)
	}, outro)
}
// don't touch below this line

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
