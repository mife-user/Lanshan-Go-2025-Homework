package main

import "fmt"

var cyjj int
var xgjj int
var cy_xgjj float64
var xg_cyjj int

func average(sum int, count int) float64 {
	cy_xgjj = float64(sum) / float64(count)
	return cy_xgjj
}
func main() {

	for i := 0; ; i++ {
		fmt.Println("请输入一个整数(输入0结束):")
		fmt.Scanln(&cyjj)
		if cyjj == 0 {
			break
		} else {
			xgjj += cyjj
			xg_cyjj++
			cyjj = 0
		}
	}
	if average(xgjj, xg_cyjj) >= 60 {
		fmt.Printf("cyjj平均成绩为：%f,成绩合格", average(xgjj, xg_cyjj))
	} else {
		fmt.Printf("xgjj平均成绩为：%f,成绩不合格", average(xgjj, xg_cyjj))
	}
}
