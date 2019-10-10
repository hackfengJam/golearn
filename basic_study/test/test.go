package main

import (
	"fmt"
	"golearn/basic_study/test/gls"
	"math/rand"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

type UserProfileInfo map[string]interface{}

func (upi UserProfileInfo) GetValue(key string) interface{} {
	keys := strings.Split(key, ".")
	return upi.getValue(keys)
}
func (upi UserProfileInfo) getValue(keys []string) interface{} {
	v := upi[keys[0]]
	if len(keys) == 1 {
		return v
	}
	vMap, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	return UserProfileInfo(vMap).getValue(keys[1:])
}

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
	/*
			// We consider a cycle to be: sweep termination, mark, mark
			// termination, and sweep. This function shouldn't return
			// until a full cycle has been completed, from beginning to
			// end. Hence, we always want to finish up the current cycle
			// and start a new one. That means:
			//
				1.在扫描终止，标记或循环N的标记终止中，等待直到标记终止N完成并过渡到扫描N。
				2.在扫描N中，帮助进行扫描N。
			At this point we can begin a full cycle N+1.
				3.通过开始扫描终止N + 1触发周期N + 1
				4.等待标记终止N + 1完成。
				5.帮助扫描N + 1，直到完成。
		   必须编写所有这些文件来处理GC可能会自行前进的事实。例如，当我们阻塞直到标记终止N时，我们可能会在周期N + 2中醒来。
	*/
	vMap := make(map[string]interface{}, 2)
	fmt.Println(UserProfileInfo(vMap).getValue([]string{"1"}))
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
	//var value interface{}

}
func mapKeyT() {
	var m = map[string]string{}
	m["a"] = "a"
	m["b"] = "b"
	fmt.Println(m)
	fmt.Println(m["a"])
	c, e := m["c"]
	fmt.Printf("c: %v", m["c"])
	fmt.Printf("c: %v, e: %v", c, e)
}

func ParseIntT() {
	fmt.Println(strconv.ParseInt("-1", 10, 64))
	fmt.Println(strconv.Atoi("-1"))
	fmt.Println(strconv.ParseInt("-1", 10, 0))

	s1, _ := strconv.ParseInt("-1", 10, 64)
	fmt.Println(reflect.TypeOf(s1))

	s2, _ := strconv.Atoi("-1")
	fmt.Println(reflect.TypeOf(s2))

	s3, _ := strconv.ParseInt("-1", 10, 0)
	fmt.Println(reflect.TypeOf(s3))

}

func TimeDurationT() {
	//a := time.Duration(1)
	a := time.Duration(-1) * time.Second
	b := time.Duration(1) * time.Second
	fmt.Println(a.Nanoseconds())
	fmt.Println(a.Hours() * 3600)
	fmt.Println(a.Seconds())
	fmt.Println(a < b)
}

func UrlEncodeT() {
	iUrl := "http://www.baidu.com/baidu哈哈/heh?a=1"
	urlList := strings.Split(iUrl, "/")

	for i := range urlList {
		urlList[i] = url.PathEscape(urlList[i])
	}
	fmt.Println(strings.Join(urlList, "/"))
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
	//reflectT()
	//ParseIntT()
	//mapKeyT()
	//UrlEncodeT()

	//var start time.Time
	//start = time.Now()
	//fmt.Printf("GoodsReviewList elapsed: %v", time.Now().Sub(start))

	var a []*int
	for i := range a {
		fmt.Println("1", i)
	}
	fmt.Println(1)
}
