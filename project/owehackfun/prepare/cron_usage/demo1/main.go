package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)

	// 秒粒度，年配置（2018-2099）
	// 某分钟（0-59），某小时（0-23），某天（1-31），某月（1-12），星期几（0-6）

	//if expr, err = cronexpr.Parse("* * * * *"); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// 每个5分钟执行1次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	// 当前时间
	now = time.Now()

	// 下次调度时间
	nextTime = expr.Next(now)
	//fmt.Println(now, nextTime)

	// 等待这个定时器超时
	//time.NewTimer(nextTime.Sub(now))
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了", nextTime)
	})

	time.Sleep(2 * time.Second)
	expr = expr

}
