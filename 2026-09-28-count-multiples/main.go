package main

import "fmt"

func countMultiples(n int, k int) int {
	contador := 0
	if k == 0 {
		return 0
	}
	for i := 1; i <= n; i++ {
		if i%k == 0 {
			contador++
		}
	}

	return contador
}

func main() {
	fmt.Println(countMultiples(10, 1))
	fmt.Println(countMultiples(10, 0))
	fmt.Println(countMultiples(10, 3))
}
