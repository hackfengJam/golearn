package main

import (
	"encoding/json"
	"fmt"
)

type PayMethod int32

const (
	PayMethod_WXPay          PayMethod = 0
	PayMethod_Invalid        PayMethod = 1
	PayMethod_QPay           PayMethod = 2
	PayMethod_AliPay         PayMethod = 3
	PayMethod_System         PayMethod = 4
	PayMethod_ApplePay       PayMethod = 6
	PayMethod_WXPayV3        PayMethod = 7
	PayMethod_AliPay_Pcredit PayMethod = 9
	PayMethod_JoinPay_WX     PayMethod = 11
	PayMethod_Xy_WxPay       PayMethod = 13
	PayMethod_BaiduPay       PayMethod = 14
	PayMethod_EcpssPay_WX    PayMethod = 17
)

// type T map[PayMethod]T1
type T map[PayMethod]interface{}

type T1 struct {
	Total int `json:"total"`
}

type T2 struct {
	Total int `json:"total"`
}

func main() {
	s := `{0: {"total": 1}, 1: {"total": 2}}`
	t := &T{}

	err := json.Unmarshal([]byte(s), t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*t)
}
