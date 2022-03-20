package main

import "fmt"

func main() {
	var (
		n       int      = 100 //학생 수
		score   [100]int = [100]int{}
		largest int
	)

	largest = score[0]
	for i := 1; i < n; i++ {
		if score[i] > largest {
			largest = score[i]
		}
	}

	fmt.Println(largest)
}
