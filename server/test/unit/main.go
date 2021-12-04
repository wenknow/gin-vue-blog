package main

import "fmt"

func main() {
	var sortArr = []int{0, 0, 0}
	var keyArr = []int{0, 0, 0}
	var list = []int{0}
	for k, v := range list {
		if v > 0 {
			for i := 0; i < 3; i++ {
				if v > sortArr[i] {
					sortArr = append(sortArr[:i], append(append([]int{v}, sortArr[i:2]...))...)
					keyArr = append(keyArr[:i], append(append([]int{k}, keyArr[i:2]...))...)
					break
				}
			}
		}
	}

	fmt.Println(sortArr)
	fmt.Println(keyArr)
}
