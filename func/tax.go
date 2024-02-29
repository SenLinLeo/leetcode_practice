package main

import (
	"fmt"
	"strconv"
)

/**
公式：
累计应纳税所得额=月薪*月数-社保公积金*月数-专项附加扣除*月数-起征点*月数
当月个税=累计应纳税所得额*年度税率-年度扣除数-当年度累计已缴纳个税
**/
func getYear() string {
	afterTax := 92820.00
	return fmt.Sprintf("%0.2f", afterTax)
}

// https://www.gerensuodeshui.cn/city/%E6%B7%B1%E5%9C%B3.html
func getMonth() string {
	tax := 669.25
	resAmount := 29400 + 200 - 2826.95 - 3945.00 - tax

	return fmt.Sprintf("%0.2f", resAmount)
}

func main() {
	startAmount := 13494.81
	fYear, _ := strconv.ParseFloat(getYear(), 64)
	fmt.Println("第一笔getYear():", getYear(), " 余额: ", startAmount+fYear)
	fAmount, _ := strconv.ParseFloat(getMonth(), 64)
	fmt.Println("第二笔getMonth():", getMonth(), " 余额: ", startAmount+fYear+fAmount)
}
