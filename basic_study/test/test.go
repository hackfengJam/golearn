package main

import (
	"fmt"
	"golearn/basic_study/test/gls"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

func mapT() {
	data := make(map[int64]string, 2)

	data[1] = "aaa"
	data[2] = "bbb"
	data[3] = "ccc"
	data[12312312] = "dddd"

	fmt.Println(data)

	for _, v := range data {
		oldV := v
		ln := len(v) - 1
		v = v[0:ln] + "0" + v[ln+1:]
		fmt.Printf("%s ", v)
		fmt.Printf("%s ", oldV)
	}

	fmt.Println()
	fmt.Println(data)

}

func arrT() {
	data := "abcdefgh"
	fmt.Println(data[:len(data)])

}

func sprintFT() {
	data := "abcdefgh"
	s := fmt.Sprintf("%s", data[len(data):len(data)])
	fmt.Println(s)

}

func set2Map(data map[string]string, k string, v string) {
	if _, ok := data[k]; !ok {
		// 不存在映射则添加映射至 data
		data[k] = v
	}
}

func mapAdT() {
	data := map[string]string{"a": "1", "b": "2", "c": "3"}
	fmt.Println(data)
	set2Map(data, "a", "2")
	set2Map(data, "d", "4")
	fmt.Println(data)
}
func mulZero() {

	str := "%0" + fmt.Sprintf("%d", 3) + "s"
	fmt.Println(fmt.Sprintf(str, "0"))
}

func sliceT() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	b := a[:1]
	fmt.Println(b)
	b = append(b, 10)
	b = append(b, 10)
	b = append(b, 10)
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(b)
	fmt.Println(cap(b))
	fmt.Printf("%p %p", unsafe.Pointer(&a[0]), b)

}

func IntT() {
	//i := 10
	//i = 2 / 3
	a := []int{1, 2, 3, 4, 5}
	//segCount := len(a) /
	fmt.Println(a[:(len(a)/3)*3])

}

func AppendSlice() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl)
	s1 := sl[:2]
	fmt.Println(s1)
	s1 = append(s1, sl[1])
	fmt.Println(s1, sl)
	fmt.Println(&s1[0], &sl[0])
}

func argsT(a, b int, c string, d string) {
	fmt.Println(a, b, c, d)
}

func ArgsT(args ...interface{}) {
	uid, orderId, seq, mime := args[0].(int), args[1].(int), args[2].(string), args[3].(string)
	argsT(uid, orderId, seq, mime)
}

func ChanT() {
	closeC := make(chan int)
	limit := 3
	for i := 0; i < limit; i++ {
		go func(index int) {
			for {
				select {
				case <-closeC:
					fmt.Printf(" %d is close ", index)
					goto END
				}
			}
		END:
		}(i)
	}
	//closeC <- 0
	//closeC <- 0
	//closeC <- 0
	close(closeC)

	time.Sleep(60 * time.Second)
}

var battle = make(chan string)

func warrior(name string, done chan struct{}) {
	select {
	case opponent := <-battle:
		fmt.Printf("%s beat %s\n", name, opponent)
	case battle <- name:
		// I lost :-(
	}
	done <- struct{}{}
}

func warriorT() {
	done := make(chan struct{})
	langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
	for _, l := range langs {
		go warrior(l, done)
	}
	for _ = range langs {
		<-done
	}
}

func waiter(i int, block, done chan struct{}) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println(i, "waiting...")
	<-block
	fmt.Println(i, "done!")
	done <- struct{}{}
}

func waiterT() {
	block, done := make(chan struct{}), make(chan struct{})
	for i := 0; i < 4; i++ {
		go waiter(i, block, done)
	}

	time.Sleep(5 * time.Second)

	close(block)
	for i := 0; i < 4; i++ {
		<-done
	}
}

func RuntimeStackT() {
	var buf = make([]byte, 64)
	var stk = buf[:runtime.Stack(buf, false)]
	print(string(stk))
}

func GoidLocalT() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			defer gls.Clean()

			defer func() {
				fmt.Printf("%d: number = %d\n", idx, gls.Get("number"))
			}()
			gls.Put("number", idx+100)
		}(i)
	}
	wg.Wait()
}

func reflectT() {
	os.OpenFile("./test.txt", os.O_RDWR, 0)
	fmt.Println(reflect.ValueOf(nil), reflect.TypeOf(nil))
	var value interface{}

}

func main() {
	//mapT()
	//arrT()
	//sprintFT()
	//mapAdT()
	//sliceT()
	//IntT()
	//v := []int{1,2,3}
	//fmt.Println(v[:0])
	//AppendSlice()
	//ArgsT(1,2,"2",5,"12")
	//ChanT()
	//warriorT()
	//waiterT()
	//RuntimeStackT()
	//fmt.Println(gls.GetGoid())
	//GoidLocalT()
	reflectT()
}
