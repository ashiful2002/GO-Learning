package main

import "fmt"

func main() {

	day := "Friday"

	// switch day {
	// case "Monday":
	// 	fmt.Printf("%s It's the start of the week.\n", day)
	// case "Tuesday":
	// 	fmt.Printf("%s It's the second day of the week.\n", day)
	// case "Wednesday":
	// 	fmt.Printf("%s It's the middle of the week.\n", day)
	// case "Thursday":
	// 	fmt.Printf("%s It's almost the weekend.\n", day)
	// case "Friday":
	// 	fmt.Printf("%s It's the last day of the workweek.\n", day)
	// case "Saturday", "Sunday":
	// 	fmt.Printf("%s It's the weekend!\n", day)
	// default:
	// 	fmt.Println("Invalid day.")
	// }

	switch {
	case day == "Friday":
		fmt.Printf("%s It's the last day of the workweek.\n", day)
	case day == "Monday":
		fmt.Printf("%s It's the start of the week.\n", day)
	case day == "Tuesday":
		fmt.Printf("%s It's the second day of the week.\n", day)
	default:
		fmt.Println("Invalid day.")
	}

}