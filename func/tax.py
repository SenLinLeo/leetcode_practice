#!/usr/bin/python
# -*- coding: UTF-8 -*-
# 新个税计算
# Fox
# 用法：python tax.py <税前月薪> <社保/公积金个人缴纳> <子女、老人、房贷月专项扣除(没有填0)>
# 示例：python tax.py 30000 4500 2000

import sys

#起征点
threshold = 5000

salary = int(sys.argv[1])
deduct1 = int(sys.argv[2])
deduct2 = int(sys.argv[3])
deduct = threshold + deduct1 + deduct2

print ("月份    应发工资        社保扣除        累计计税收入    累计免税扣除    税率    纳税    累计纳税        税后")

amount = 0;
for i in range(1, 13):
        taxable = salary * i - deduct * i
        if taxable < 36000:
                r = 0.03
                rd = 0
        elif taxable >= 36000 and taxable < 144000:
                r = 0.1
                rd = 2520
        elif taxable >= 144000 and taxable < 300000:
                r = 0.20
                rd = 16920
        elif taxable >= 300000 and taxable < 420000:
                r = 0.25
                rd = 31920
        elif taxable >= 420000 and taxable < 660000:
                r = 0.30
                rd = 52920
        elif taxable >= 660000 and taxable < 960000:
                r = 0.35
                rd = 85920
        else: # if taxable >= 960000:
                r = 0.45
                rd = 181920

        tax = int(taxable * r - rd - amount)
        amount += tax

        aftax = salary - tax - deduct1

        print ("%s      %s              %s              %s              %s              %s      %s      %s              %s" % (i, salary, deduct1, taxable, deduct*i, r, tax, amount, aftax))

#python tax.py 29600 6772 0