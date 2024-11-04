package main

import "fmt"

func main() {
	demo_break_no_label()
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

func demo_four() {
	var sum int
	for i, j, k := 0, 1, 2; (i < 20) && (j < 10) && (k < 30); i, j, k = i+1, j+1, k+5 {
		sum += i + j + k
	}
	fmt.Println(sum)
}

// continue loop  结束当前循环，继续下一次循环
func demo_five() {
	var sum int
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
loop:
	for _, v := range data {
		if v%2 == 0 {
			continue loop
		}
		sum += v
	}

	fmt.Println(sum)
}

// break loop 结束循环
func demo_six() {
	var sum int
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
loop:
	for i := 0; i < 10; i++ {
		if data[i]%2 == 0 {
			break loop
		}
		sum += data[i]
	}
	fmt.Println(sum)
}

// 嵌套循环语句中，被用于跳转到外层循环并继续执行外层循环语句的下一个迭代，比如下面这段代码：
func demo_seven() {
	var sl = [][]int{{1, 34, 26, 35, 78}, {3, 45, 13, 24, 99}, {101, 13, 38, 7, 127}, {54, 27, 40, 83, 81}}
outerloop:
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == 13 {
				fmt.Printf("found 13 at [%d, %d]\n", i, j)
				continue outerloop
			}
		}
	}
}

// 通过标签跳转到指定的循环
// Go 也 break 语句增加了对 label 的支持。而且，和前面 continue 语句一样，如果遇到嵌套循环，break 要想跳出外层循环，用不带 label 的 break 是不够，因为不带 label 的 break 仅能跳出其所在的最内层循环。要想实现外层循环的跳出，我们还需给 break 加上 label
func demo_break_label() {
	var gold = 38
	var sl = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}
outerloop2:
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == gold {
				fmt.Printf("found gold at [%d, %d]\n", i, j)
				break outerloop2
			}
		}
		fmt.Println("continue outer loop")
	}
}

func demo_break_no_label() {
	var gold = 38
	var sl = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == gold {
				fmt.Printf("found gold at [%d, %d]\n", i, j)
				break
			}
		}
		fmt.Println("continue outer loop")
	}
}
