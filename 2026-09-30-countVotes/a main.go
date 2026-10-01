package main

import "fmt"

func countVotes(votes []string) map[string]int {
	count := map[string]int{}
	for _, name := range votes {
		count[name]++
	}
	return count
}

func main() {
	list := []string{"ana", "bia", "ana"}
	result := countVotes(list)
	fmt.Println(result)
}
