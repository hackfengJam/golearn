package fib

func Fibnoacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// chan 实现迭代器
func FibnoacciChan() (c chan int) {

	// init
	c = make(chan int, 0)
	a, b := 0, 1

	cnt := 1

	go func() {
		for {
			// 终止条件
			if cnt > 10 {
				break
			}
			// 迭代
			cnt++

			// 获取输出，具体逻辑
			a, b = b, a+b

			// 输出
			c <- a
		}
		close(c)
	}()
	return
}
