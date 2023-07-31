package main

import "fmt"

func adder() func(int) int {
	sum := 0 // 戻り値の関数が sum に（環境に）バインドされる
	return func(x int) int {
		sum += x
		return sum
	}
}

func practiceFunctionClosures() {
	fmt.Println("----- practice function closures -----")
	pos, neg := adder(), adder() // pos, neg がそれぞれの sum にバインドされる
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/* 出力
----- practice function closures -----
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
*/