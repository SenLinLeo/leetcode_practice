package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
)

/**
公式：
累计应纳税所得额=月薪*月数-社保公积金*月数-专项附加扣除*月数-起征点*月数
当月个税=累计应纳税所得额*年度税率-年度扣除数-当年度累计已缴纳个税
**/
func getYear() float64 {
	afterTax := 92820.00
	return afterTax
}

// https://www.gerensuodeshui.cn/city/%E6%B7%B1%E5%9C%B3.html
func getMonth() float64 {
	tax := 455.11 + 614.35
	resAmount := 29400 + 200 - 2826.95 - 3945.00 - tax

	return resAmount
}

func main() {
	startAmount := 1494.02
	fYear := humanize.Commaf(getYear())
	fMonth := humanize.Commaf(getMonth())
	fmt.Println("第1笔Mon:", fMonth, "\t余额: ", humanize.Commaf(startAmount+getYear()+getMonth()))
	fmt.Println("第2笔Yea:", fYear, "\t余额: ", humanize.Commaf(startAmount+getYear()))
	/**
	第1笔Mon: 21,758.59     余额:  116,072.61
	第2笔Yea: 92,820        余额:  94,314.02
	**/
	fmt.Println(humanize.Commaf(102900.00))
	fmt.Println(humanize.Commaf(102900.00 - 29400.00 - 29400.00))
}
