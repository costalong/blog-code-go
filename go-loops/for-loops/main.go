package main

import "fmt"

func main() {
	demo_three()
}

func demo_one() {
	// Simple for loop
	for i := 0; i < 5; i++ {
		println(i)
	}
}

// 为偶数的时候跳出
func demo_two() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}
}

// 为偶数的时候跳出循环
func demo_three() {
	var end int
	for i := 1; i < 10; i++ {
		end = i
		if i%2 == 0 {
			break
		}
		fmt.Println(i)
	}
	println(end)
}
