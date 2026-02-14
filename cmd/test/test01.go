package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	var count int
	for N > 0 {
		var S string
		S = strconv.Itoa(N)
		var tmp int
		for i := 0; i < len(S); i++ {
			val, _ := strconv.Atoi(string(S[i]))
			tmp += val
		}
		if tmp == K {
			count++
		}
		N--
	}
	fmt.Println(count)
}
