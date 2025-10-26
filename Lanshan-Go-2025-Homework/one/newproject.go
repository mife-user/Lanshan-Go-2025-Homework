package main

import "fmt"

const pi float32 = 3.14

var a float32 = 5
var Plus = 0
var cyjj = 0
var n int

type ruan func(float32, float32) float32

func in_put(m float32, n float32) float32 {
	fmt.Println("结果为：", m*n*n)
	return 0
}
func j_c(k int) int {
	if k == 1 {
		return 1
	} else {
		cyjj += j_c(k-1) + k
	}
	return cyjj
}
func main() {
	fmt.Println("hellow mife")
	var yu ruan = in_put
	yu(pi, a)
	for i := 1; i < 1001; i++ {
		Plus += i
	}
	fmt.Println("1~1000的和为：")
	fmt.Println(Plus)
	fmt.Println("请输入你要求的阶乘：")
	fmt.Scanln(&n)
	fmt.Println("结果为L:")
	fmt.Println(j_c(n))

}
