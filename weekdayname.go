package main

import "fmt"

func weekdayname(day int) string {
	switch day {
	case 1:
		return "monday"
	case 2:
		return "tuesday"
	case 3:
		return "wednesday"
	case 4:
		return "thursday"
	case 5:
		return "friday"
	case 6:
		return "saturday"
	case 7:
		return "sunday"
	default:
		return "invalid"

	}

}
func main() {
	fmt.Println(weekdayname(5))
	fmt.Println(weekdayname(7))
	fmt.Println(weekdayname(2))

}
