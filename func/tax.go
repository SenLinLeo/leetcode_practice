package main

import (
	"fmt"
	"strconv"
	"github.com/dustin/go-humanize"
)

/**
公式：
累计应纳税所得额=月薪*月数-社保公积金*月数-专项附加扣除*月数-起征点*月数
当月个税=累计应纳税所得额*年度税率-年度扣除数-当年度累计已缴纳个税
**/
func getYear() string {
	afterTax := 92820.00
	return fmt.Sprintf("%v", humanize.Commaf(afterTax))
}

// https://www.gerensuodeshui.cn/city/%E6%B7%B1%E5%9C%B3.html
func getMonth() float64 {
	tax := 455.11 + 614.35
	resAmount := 29400 + 200 - 2826.95 - 3945.00 - tax

	return resAmount
}

func main() {
	startAmount := 13494.81
	fYear, _ := strconv.ParseFloat(getYear(), 64)
	fmt.Println("第1笔getMonth():", humanize.Commaf(getMonth()), " 余额: ", humanize.Commaf(startAmount+fYear))
	fmt.Println("第2笔getYear():", getYear(), " 余额: ", humanize.Commaf(startAmount+fYear+getMonth()))
	/**
	第1笔getMonth(): 21,758.59  余额:  13,494.81
	第2笔getYear(): 92,820  余额:  35,253.4
	**/
}
