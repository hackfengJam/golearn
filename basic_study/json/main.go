package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Student struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type Class struct {
	Id       int       `json:"id,omitempty"`
	Students []Student `json:"students,omitempty"`
}

type School struct {
	Id      int             `json:"id"`
	Classes map[int64]Class `json:"classes"`
}

func basicJSON() {
	var stu Student

	sJson := `{
	"id": 1,
	"name": "zhangsan",
	"age": 3
}`
	if err := json.Unmarshal([]byte(sJson), &stu); err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu.Id)
}
func arrJSON() {
	var cls Class

	sJson := `{
	"id": 2,
	"students": [
{
	"id": 1,
	"name": "zhangsan",
	"age": 3
},
{
	"id": 2,
	"name": "lisi",
	"age": 30
},
{
	"id": 3,
	"name": "xiaoming",
	"age": 30
}
]
}`

	if err := json.Unmarshal([]byte(sJson), &cls); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(cls.Id)
	fmt.Println(cls.Students)
}

func mapJSON() {
	var cls School

	sJson := `{
	"id": 2,
	"classes": {
		"1" : {
			"id": 2,
			"students": [
		{
			"id": 1,
			"name": "zhangsan",
			"age": 3
		},
		{
			"id": 2,
			"name": "lisi",
			"age": 30
		},
		{
			"id": 3,
			"name": "xiaoming",
			"age": 30
		}]}, 
		"2" : {
			"id": 2,
			"students": [
		{
			"id": 1,
			"name": "zhangsan2",
			"age": 3
		},
		{
			"id": 2,
			"name": "lisi2",
			"age": 30
		},
		{
			"id": 3,
			"name": "xiaoming2",
			"age": 30
		}
		]}
		}
}`

	if err := json.Unmarshal([]byte(sJson), &cls); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(cls.Id)
	fmt.Println(cls.Classes)
}
func nilJSON() {

	var stu Student

	sJson := `{
	"id": 1,
	"name": "zhangsan",
	"age": 3
}`

	if err := json.Unmarshal([]byte(sJson), &stu); err != nil {
		fmt.Println(err)
	}
	stu.Id = 0

	fmt.Println(stu.Id)

	str, _ := json.Marshal(stu)
	fmt.Println(string(str))
}

func nilMapJSON() {
	var cls Class

	sJson := `{
	"id": 2,
	"students": [
{
	"id": 1,
	"name": "zhangsan",
	"age": 3
},
{
	"id": 2,
	"name": "lisi",
	"age": 30
},
{
	"id": 3,
	"name": "xiaoming",
	"age": 30
}
]
}`

	if err := json.Unmarshal([]byte(sJson), &cls); err != nil {
		fmt.Println(err)
	}

	cls.Id = 0
	cls.Students = make([]Student, 0)

	fmt.Println(cls)
	fmt.Println(cls.Students)

	str, _ := json.Marshal(cls)
	fmt.Println(string(str))
}

func typeJSON() {
	var stu Student

	sJson := `{
		"id": "1",
		"name": "zhangsan",
		"age": 3
	}`
	if err := json.Unmarshal([]byte(sJson), &stu); err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)
}

func main() {
	//arrJSON()
	//mapJSON()
	//nilJSON()
	//nilMapJSON()
	//typeJSON()

	x := 0.01
	fmt.Println(fmt.Sprintf("%v推币抵扣%v", x, x*1.0/100))

	fmt.Println(strings.Join([]string{"activity", "uniqueCode"}, "."))
}

/*

f(i, c) = f(i - 1, c)
        = f(i - 1, c - v(i)) + w(i)

重量 c，满足 wl[i] 可以装重 w[i]； 求最重。

f(i, c) = f(i - 1, c)


1000-300 1000-200 1000-100 900-400 800-300 700-100

     100 200 300 400 500 600 700 800 900 1000 1100 1200 1300
      0   0   0   0   0   0   0   0   0   300  300  300  300
      0   0   0   0   0   0   0   0   0   300  300  500  500
      0   0   0   0   0   0   0   0   0   200  300  500  500
      0   0   0   0   0   0   0   0   0   300  400  400  400
      //0   0   0   0   0   0   0   0   0   100  100  100  300
      0   0   0   0   0   0   0   0   0   300  400  400  400

1000-300 1000-200 {1000-100 900-400 800-300} 700-100

     100 200 300 400 500 600 700 800 900 1000 1100 1200 1300
      0   0   0   0   0   0   0   0   0   300  300  300  300
      //0   0   0   0   0   0   0   0   0   0   0   0   0
      0   0   0   0   0   0   0   0   0   300  300  300  300
      //0   0   0   0   0   0   0   0   0   200  200  200  200
      0   0   0   0   0   0   0   0   0   300  400  400  400
      //0   0   0   0   0   0   0   0   0   100  100  100  300
      0   0   0   0   0   0   0   0   0   300  400  400  400

3^100

     100 200 300 400 500 600 700 800 900 1000 1100 1200 1300
      0   0   0   0   0   0   0   0   0   300  300  300  300
      //0   0   0   0   0   0   0   0   0   0   0   0   0
      0   0   0   0   0   0   0   0   0   300  300  300  300
      //0   0   0   0   0   0   0   0   0   200  200  200  200
      0   0   0   0   0   0   0   0   0   300  400  400  400
      //0   0   0   0   0   0   0   0   0   100  100  100  300
      0   0   0   0   0   0   0   0   0   300  400  400  400

100 - 100   200
50  - 150

100 - 200   300
50  - 250

350 - 100

100 - 100
100 - 200

50  - 250

*/
/*
算法：

首先判断一个分组当中的一件物品，同01背包一样，此物品存在两种状态，取与不取，若取此物品，则继续判断下一组的第一件物品，
若不取此物品，则继续判断本组下一件物品，若该物品为本组最后一件物品，则判断下一组。
也就是说设f[k][v]表示前k组物品花费费用v能取得的最大权值，则有：f[k][v]=max{f[k-1][v],f[k-1][v-c[i]]+w[i]|物品i属于组k}。
使用一维数组的伪代码如下：

————————————————

for 所有的组k
	for v=V..0
		for 所有的i属于组k
			f[v]=max{f[v],f[v-c[i]]+w[i]}



优惠券结算问题

一、购物经常碰到的优惠券结算问题；现有 平台券 和 店铺券；如下描述：

平台：[优惠券 1 优惠券 2 优惠券 3]
店铺：
	- 店铺 1：[优惠券 1 优惠券 2 优惠券 3]
	- 店铺 2：[优惠券 1 优惠券 2 优惠券 3]
	- 店铺 3：[优惠券 1 优惠券 2 优惠券 3]
	...

二、一些变量定义：

1. 账单cost[i]，表示用户在第i个店铺所消费金额
如：
[200, 300]

2. coupon[i][j] 表示第i个店铺第j个 可用优惠券。因为是满减优惠券，coupon[i][j]是存在两个元素的一维数组指：满 coupon[i][j][0] 减 coupon[i][j][1]。
如：
[
[[200,100], [150,50]],
[[300,150], [200,50]]
]
ps: [200,100] 即 满200减100

3. platform[k] 代表平台下 第k个可用优惠券，platform[k][0] 减 platform[k][1]
如：
[[300,100], [200,50]]

三、约束与问题

1. 约束：
平台在可用的优惠券中至多选择一个，店铺下每个店铺至多选择一个可用优惠券（均为满减优惠券）。

2. 问题
给出一个账单 cost，coupon，platform；找出一组减免最多的优惠券组合。
注：
结果 为 ret 一位数组；长度为 len(coupon) + 1；即为所有 店铺的选择以及对平台的选择；
- ret[0...len(coupon)] 意指给每个店铺选择的优惠券；ret[i] 代表 第i个店铺选择的 可用优惠券索引号
- ret[len(coupon)] 指给平台选择的优惠券
- -1 代表不使用相应优惠券

四：关于示例解释
示例：
- cost
[200, 300]

- coupon
[
[[200,100], [150,50]],
[[300,150], [200,50]]
]

- platform
[[300,100], [200,50]]

结果
ret = [1, 0, 0]
ps：
- 店铺 0：选择 满150减50
- 店铺 1：选择 满300减150
- 平台  ：选择 满300减100

订单在每家店铺的消费金额为 [200, 300]；因此对于结果 [1, 0, 0]，解释为：共消费500元，最大可减免300元。

leetcode 198 打家劫舍
你是一个专业的小偷，打算洗劫一条街的所有房子。
每一个房子里都有不同价值的宝物；
但是，如果你选择偷窃连续两栋的房子，就会出发报警系统。编程求出你最多可以偷窃价值多少的宝物？

暴力解法：检查所有房子的组合，对每一个组合，检查是否有相邻的房子，如果没有，记录其价值。找最大值。O((2^n)*n)

其中对状态的定义：
考虑偷取 [x...n-1] 范围里的房子 （函数的定义）

根据对状态的定义，决定状态的转移：
f(0) = max{ v(0) + f(2), v(1) + f(3), v(2) + f(4), ... , v(n-3) + f(n-1), v(n-2), v(n-1) }


其中对状态的定义：
考虑偷取 [0...x] 范围里的房子 （函数的定义）

根据对状态的定义，决定状态的转移：
f(n-1) = max{ v(n-1) + f(n-3), v(n-2) + f(n-4), ... , v(2) + f(0), v(1), v(0) }

*/
