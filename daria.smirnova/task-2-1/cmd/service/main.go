package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Error while reading input:", err)

		return
	}

	for range departments {
		var employees int

		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("Error while reading input:", err)

			return
		}

		low, high := minTemp, maxTemp

		for range employees {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("Error while reading input:", err)

				return
			}

			switch operator {
			case ">=":
				low = max(low, temp)
			case "<=":
				high = min(high, temp)
			}

			fmt.Println(low)
		}
	}
}
