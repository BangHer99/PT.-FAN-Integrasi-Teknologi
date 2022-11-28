package main

import "fmt"

// SOAL NOMOR 1 DONE

func Sock(ar []int) int {
	count := 0
	check := make(map[int]bool)

	for _, num := range ar {
		if val, exist := check[num]; exist {
			if val {
				count++
				check[num] = false
				continue
			}
		}
		check[num] = true
	}
	return int(count)
}

func main() {
	fmt.Println(Sock([]int{10, 20, 20, 10, 10, 30, 50, 10, 20}))
	fmt.Println(Sock([]int{6, 5, 2, 3, 5, 2, 2, 1, 1, 5, 1, 3, 3, 3, 5}))
	fmt.Println(Sock([]int{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}))
}
