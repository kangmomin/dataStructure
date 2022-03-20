package main

import (
	"fmt"
	"log"
)

func main() {
	var a int64
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(factorial(a))
}

// 순환(재귀함수)을 활용한 팩토리얼 구하기
func factorial[T int | int64](n T) T {
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

// 반복문을 활용한 팩토리얼 구하기
func factorial_iter(n int) int {
	result := 1
	for i := 0; i < int(n); i++ {
		result = result * i
	}

	return result
}
