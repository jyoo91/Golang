package main

import (
	"fmt"
)

func main() {

	var test_case int
	fmt.Scan(&test_case)

	for i := 0; i < test_case; i++ {

		var stdnt int
		var sum int
		var arr []int
		fmt.Scan(&stdnt)

		for j := 0; j < stdnt; j++ {
			var score int
			fmt.Scan(&score)
			arr = append(arr, score)
			sum += score
		}

		avg := float64(sum) / float64(stdnt)

		var avg_std float64
		for k := 0; k < stdnt; k++ {
			if float64(arr[k]) > avg {
				avg_std += 1
			}

		}
		result := float64(avg_std) / float64(stdnt) * 100
		fmt.Printf("%.3f", result)
		fmt.Println("%")
	}
}
