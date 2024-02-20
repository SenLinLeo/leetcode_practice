package main

import "fmt"

/**
公式：
累计应纳税所得额=月薪*月数-社保公积金*月数-专项附加扣除*月数-起征点*月数
当月个税=累计应纳税所得额*年度税率-年度扣除数-当年度累计已缴纳个税
**/
func getYear() string {
	totalYear := 29400 * 3.5
	afterTax := totalYear - totalYear*0.1
	return fmt.Sprintf("%0.2f", afterTax)
}

func getMonth() string {
	// 调整项
	otherAmount := 10340.00
	otherTax := otherAmount * 0.03
	resTax := 1739.37 - otherTax
	resAmount := 29400 + 200 - resTax - 6771.95

	return fmt.Sprintf("%0.2f", resAmount)
}

func main() {
	fmt.Println("getYear():", getYear(), "--", 13494.81+92610.00)            // 92610.00
	fmt.Println("getMonth():", getMonth(), "--", 13494.81+92610.00+21398.88) // 21398.88
	// 13494.81
	fmt.Println(48414.81 - 34920.00)
}
