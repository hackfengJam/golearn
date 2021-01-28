package main

import (
	"bufio"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"golearn/basic_study/test/gls"
	"golearn/basic_study/test/sdk"
	// "golearn/basic_study/test/util/simpleaes"
	"math"
	"math/rand"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"
	"unsafe"
)

func main() {

	// fmt.Printf("main: %d\n", name)

	// mapT()
	// arrT()
	// sprintFT()
	// mapAdT()
	// sliceT()
	// IntT()
	// v := []int{1,2,3}
	// fmt.Println(v[:0])
	// AppendSlice()
	// ArgsT(1,2,"2",5,"12")
	// ChanT()
	// warriorT()
	// waiterT()
	// RuntimeStackT()
	// fmt.Println(gls.GetGoid())
	// GoidLocalT()
	// reflectT()
	// ParseIntT()
	// mapKeyT()
	// UrlEncodeT()
	// UrlT()
	// var start time.Time
	// start = time.Now()
	// fmt.Printf("GoodsReviewList elapsed: %v", time.Now().Sub(start))
	// var a []*int
	// for i := range a {
	//	fmt.Println("1", i)
	// }
	// fmt.Println(1)
	// JsonMapInterfaceT()
	// JsonMapStructT()
	// regexpT()
	// bitOrT()
	// shiftRightLogicalT()
	// gRpcDebugT()
	// Authorization4SignT()
	// NewClientT()
	// SortSliceT()
	// DateFormatT()
	// MapIfT()
	// NewT()
	// ContextT()
	// StringsUtils()
	// InterfaceJsonT()
	// ReflectFieldByNameT()
	// JsonT()
	// mulChar()
	// CeilT()
	// YamlToJSONT()
	// JsonToYamlT()
	// SwitchT()
	// TimeT()
	// _ = ErrorT()
	// timeT()
	// deferT()
	// GoschedT()
	// TimeAddT()
	// deferPanicT()
	// fmt.Println(fmt.Sprintf("%d", IEOthers))
	// sliceAppendT()
	// AtomicT()
	// UnicodeT()
	// timeLocalT()
	// fileT()
	// NewClientT()
	// fmt.Println(string([]byte{}))
	// copyT()
	// _ = aseT()
	// timeLocT()
	// nilInCallT()
	// funcT()
	// JsonPointT()
	// var a interface{}
	// if a == nil {
	// 	fmt.Println(1)
	// }
	// MapNilT(nil)
	// GoRoutineT()
	// deferArgsT()

	// WaitGroupT()
	// GoschedT()
	// ExampleUnmarshal()
	// SwitchPrefT()

	// MapPrefT()

	// ja := JAA{a: 2, B: time.Now()}
	// //
	// // j := JBB{ja}
	// // j.a = 1
	// // j.JAAT()
	// jaB, err := json.Marshal(ja)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(string(jaB))
	// jsonStringT()
	// tickerT()
	// varNilT()
	// fmtPointPrintT()
	// deferErrorT()
	// SwitchOrderT()

	// timeParseInLocationT()
	// mutexT()

	// chanGoT()

	// ListNodeT()

	// deferReturnT()

	// fmt.Printf("haha%shaha\n", strings.TrimSpace(" a b c  "))
	// fmt.Printf("haha%shaha\n", strings.TrimRight(" a b c  ccc", "c"))
	// fmt.Printf("haha%shaha\n", strings.TrimSuffix(" a b c  ccc", "c"))
	// TimeStr2TimeT()

	// timeSubT(2020, 12, 14)

	// fmt.Println(strings.TrimPrefix("aaa", "a"))

	// jsonMarshalT()

	// fmt.Println(time.Unix(0, 0))
	// stringToDateT()

	// now, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	// fmt.Println(err)
	// fmt.Println(now)

	return
}

type TripleTuple struct {
	Key       string
	Operation string
	Value     interface{}
}

func jsonMarshalT() {
	conditionIf := [][]interface{}{}
	ttList := []TripleTuple{
		{
			Key:       "a",
			Operation: "=",
			Value:     "1",
		},
		{
			Key:       "b",
			Operation: ">=",
			Value:     2,
		},
		{
			Key:       "c",
			Operation: "<=",
			Value:     3,
		},
	}
	for _, tt := range ttList {
		conditionIf = append(conditionIf, []interface{}{tt.Key, tt.Operation, tt.Value})
	}
	conditionJSON, err := json.Marshal(conditionIf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("conditionJSON: %s\n", string(conditionJSON))
	return
}

func timeSubT(year int, month time.Month, day int) {
	t := time.Date(year, month, day, 0, 0, 0, 0, BeijingLocation)
	fmt.Println(math.Ceil(t.Sub(time.Now()).Hours() / 24))
}

var BeijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

func TimeStr2TimeT() {
	layoutISO := "2006-01-02"
	layoutISO = "2006-01-02T15:04:05Z07:00"
	layoutISO = "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(layoutISO, "2019-04-02 19:26:24.789", BeijingLocation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
}

func deferReturnT2() (int, bool) {
	fmt.Printf("deferReturnT2: %p\n", deferReturnT2)
	return 1, false
}

func deferReturnT1() (int, bool) {
	defer func() {
		fmt.Printf("deferReturnT1: %p\n", deferReturnT1)
	}()
	return deferReturnT2()
}

func deferReturnT() {
	deferReturnT1()
}

type ListNode struct {
	Val  int32
	Val2 int32
	Next *ListNode
}

// recursive type
// type ListNode struct {
// 	Val  int
// 	Next ListNode
// }

func ListNodeT1() {
	node := ListNode{}

	// temp := node
	var temp = node

	temp.Next = &ListNode{}

	fmt.Println("ListNodeT1")
	fmt.Printf("ListNodeT1 %p\n", &node)
	fmt.Println(node)
	fmt.Printf("ListNodeT1 %p\n", &temp)
	fmt.Println(temp)
	fmt.Println(node == temp)
	fmt.Println(&node == &temp)
}

func ListNodeT2() {
	node := &ListNode{}
	// temp := node
	var temp = node

	temp.Next = &ListNode{}

	fmt.Println("ListNodeT2")
	fmt.Printf("ListNodeT2 %p\n", &node)
	fmt.Println(node)
	fmt.Printf("ListNodeT2 %p\n", &temp)
	fmt.Println(temp)
	fmt.Println(node == temp)
	fmt.Println(&node == &temp)
}

func ListNodeT() {
	ListNodeT1()
	ListNodeT2()
}
func chanGoT1(a chan int) {
	a <- 1
	// a <- 2
	// a <- 3
	// a <- 4
}
func chanGoT2(chanName string, a chan int) {
	for v := range a {
		fmt.Printf("%s, %d\n", chanName, v)
	}
}
func chanGoT() {
	ch := make(chan int, 1)
	chanGoT1(ch)
	// go chanGoT1(ch)
	close(ch)
	// defer close(ch)
	// go chanGoT2("ch1", ch)
	go chanGoT2("ch2", ch)
	time.Sleep(time.Second * 100)
}

const (
	iotaA = 1 << iota
	iotaB
	iotaC = iota
	iotaD
	iotaE = 1 << iota
	iotaF
	iotaG
)

func mutexT() {
	fmt.Printf("iotaA -> %d\n", iotaA)
	fmt.Printf("iotaB -> %d\n", iotaB)
	fmt.Printf("iotaC -> %d\n", iotaC)
	fmt.Printf("iotaD -> %d\n", iotaD)
	fmt.Printf("iotaE -> %d\n", iotaE)
	fmt.Printf("iotaF -> %d\n", iotaF)
	fmt.Printf("iotaG -> %d\n", iotaG)
	fmt.Println(1 | 2)
	fmt.Println(1 & (1 | 2))
}

func timeParseInLocationT() {
	s := "2020-08-13 13:56:41"
	transferTime, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.FixedZone("Asia/Shanghai", 8*60*60))
	if err != nil {
		return
	}
	fmt.Println(transferTime)
}

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
func mulChar() {
	fmt.Println(strings.Repeat("*", 3))
}

func sliceCapT() {
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

func sliceAppendT() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	// a = append(a, b...)
	var actIdList []int
	actIdList = append(actIdList, a...)
	actIdList = append(actIdList, b...)
	fmt.Println(a)
	fmt.Println(actIdList)
}

func IntT() {
	// i := 10
	// i = 2 / 3
	a := []int{1, 2, 3, 4, 5}
	// segCount := len(a) /
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
	// closeC <- 0
	// closeC <- 0
	// closeC <- 0
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
	// var value interface{}

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
	// a := time.Duration(1)
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

func UrlT() {
	u := url.Values{}

	var data map[string]string
	jsonStr := `
{alipay_trade_query_response":
  {"code":"10000","msg":"Success","buyer_logon_id":"793***@qq.com","buyer_pay_amount":"0.00","buyer_user_id":"2088902521835964","invoice_amount":"0.00","out_trade_no":"t-8072603435828035624","point_amount":"0.00","receipt_amount":"0.00","send_pay_date":"2020-01-06 20:23:21","total_amount":"0.01","trade_no":"2020010622001435960568692210","trade_status":"TRADE_SUCCESS"},

  "sign":"O55VLcHhBm9zNk3rqfwlUCLB1qwy8E9DmTdBukGz8+gFkpc7SNSbMh9E2/3BT7FtVX2zEtzgHGFsEcEWjHMS2568wRepPDX+rcMNYKwha+X1Zz44FUPkH03edQ2E2aj+I4vBiJsDSkErcrAwJS/7A2zQKDx8nE7kzonnrKYfgcpL8ntS3pl7ZUTjoLvdAUtkhoSyAOo463sqZAIo1LdQZo3G6S9maJMdFmDXGUpvwNHzNqxegl1rfta8IC6PPfl2aMoSAAP6qr59tH5KiN5c3t7smE3a94MhPuApORRm/vM6mzW87RNl5l37JAVqKuXIcD5tKqbLIOksYYMb+gFVqA=="}`
	jsonStr = `{"code":"10000","msg":"Success","buyer_logon_id":"793***@qq.com","buyer_pay_amount":"0.00","buyer_user_id":"2088902521835964","invoice_amount":"0.00","out_trade_no":"t-8072603435828035624","point_amount":"0.00","receipt_amount":"0.00","send_pay_date":"2020-01-06 20:23:21","total_amount":"0.01","trade_no":"2020010622001435960568692210","trade_status":"TRADE_SUCCESS"}`

	fmt.Println(data)
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	for k, v := range data {
		u.Add(k, v)
	}

	fmt.Println(u.Encode())
}

func mapModifyT(data interface{}) {
	if v, ok := data.(map[string]interface{}); !ok {
		return
	} else {
		v["a"] = "haha"
	}
	return
}

func JsonMapInterfaceT() {
	var data map[string]interface{}
	jsonStr := `{"a": 1, "b": "abcd", "c": 3}`

	fmt.Println(data)
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	mapModifyT(data)
	fmt.Println(data)

	jsonByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonByte))

}

func JsonMapStructT() {
	var data map[string]bool
	jsonStr := `{"a": 1, "b": true, "c": 3}`

	fmt.Println(data)
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	mapModifyT(data)
	fmt.Println(data)

	jsonByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonByte))

}

var getRegexp = regexp.MustCompile("^get.(.+?)(?:\\{([^\\}]+)\\})?(?:\\[(\\d+)\\])?$")

func VariableGetCreator(name string) {
	if ma := getRegexp.FindStringSubmatch(name); len(ma) == 4 {
		fmt.Println(ma)
	}
}

func regexpT() {
	VariableGetCreator("get._vars{mall_id}")
}

func bitOrT() {
	fmt.Println(1 | 2 | 4)
}

func gRpcDebugT() {
	/*
			./bin/grpcdebug -data='{"buyer_id":35899596,"malls":[{"mall_id":100003,"auto_select_mall_coupon":true,"goods":[{"goods_id":11160,"goods_sku_id":40818,"remain_price":5000,"can_use_promotion_id":[520,535],"goods_quantity":1,"goods_unit_price":5000},{"goods_id":12544,"goods_sku_id":81895,"remain_price":1400,"can_use_promotion_id":[535],"goods_quantity":1,"goods_unit_price":1400}]}],"auto_select_system_coupon":true,"plt":30,"selected_system_promotion_id":[535],"promotions":[{"promotion_id":520,"link":"/x/promotion.html?activity_id=520","label":"跨店满减","name":"满50减15,满80减35,满100减50","use_start_time":1569730200,"use_end_time":1602086400,"promotion_type":1,"rules":[{"condition_type":1,"condition_value":10000,"discount_type":1,"discount_value":5000,"mall_cost_type":1,"mall_cost_value":1500},{"condition_type":1,"condition_value":8000,"discount_type":1,"discount_value":3500,"mall_cost_type":1,"mall_cost_value":3500},{"condition_type":1,"condition_value":5000,"discount_type":1,"discount_value":1500,"mall_cost_type":1,"mall_cost_value":4900}]},{"promotion_id":535,"link":"/x/promotion.html?activity_id=535","label":"满2免1","name":"满2免1","use_start_time":1571301600,"use_end_time":1571500800,"promotion_type":1,"rules":[{"condition_type":2,"condition_value":2,"discount_type":2,"discount_value":1,"system_cost_value":10}]}],"last_selected_type":1}' -addr=127.0.0.1:9201 -method=/ptcoupon.Service/GetOrderDiscount -ctx='{"uid":1,"x-reqid":"abcdefg"}'


		./bin/grpcdebug -data='{"uids": [1094505]}' -addr=192.168.18.59:9701 -method=/ptprofile.Service/SyncUsersProfile


		./bin/grpcdebug -data='{"contract_status": 4}' -addr=127.0.0.1:8001 -method=/pay.Service/TestGrpc
	*/
}

func shiftRightLogicalT() {
	s := 1
	s += 1 << 1
	fmt.Println(s)
}

// mt :xxxx -> xxx
func Sign2Authorization(prefix string, sign string) string {
	return strings.Join([]string{
		prefix, sign,
	}, " ")
}

// xxxx -> mt :xxxxx
func Authorization2Sign(prefix string, authorization string) string {
	return strings.TrimLeft(authorization, prefix+" ")
}

func Authorization4SignT() {
	prefix := "hackfun"
	authorization := "hackfun test"
	sign := Authorization2Sign(prefix, authorization)
	fmt.Println(sign)
	fmt.Println(Sign2Authorization(prefix, sign))
}

type Entity struct {
	A int
}

func SortSliceT() {
	testCase := []Entity{
		{A: 1},
		{A: 3},
		{A: 2},
	}
	fmt.Println(testCase)
	sort.Slice(testCase, func(i, j int) bool {
		return testCase[i].A > testCase[j].A
	})
	fmt.Println(testCase)

}

type Clazz struct {
}

func NewObj() {
	var obj *Clazz
	if obj == nil {
		fmt.Println()
	}
}

const dateFormat = "2006-01-02"

func DateFormatT() {
	now := time.Now()
	fmt.Println(now.Format(dateFormat))
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(now.UTC())
	fmt.Println(now.UTC().In(local))
	fmt.Println(now.IsZero())
	fmt.Println(now)
	fmt.Println(now.In(local))
}

func MapIfT() {
	ids := []uint64{1, 2, 3, 4}
	var idsMap map[uint64]bool
	idsMap = make(map[uint64]bool, len(ids))
	for _, id := range ids {
		idsMap[id] = true
	}

	if idsMap[1] {
		fmt.Println(idsMap[1])
	}

	idsMap[2] = false
	if idsMap[2] {
		fmt.Println(idsMap[2])
	}

	if idsMap[6] {
		fmt.Println("idsMap[6]...")
	}
}

type newT struct {
}

func (n *newT) NewTFunc1() {
	fmt.Println("NewTFunc1")
}
func (n *newT) NewTFunc2() {
	fmt.Println("NewTFunc2")
}
func NewT() {
	var nS newT
	var nSC *newT
	nS.NewTFunc1()
	nSC.NewTFunc1()
}

func ContextT() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	c := context.Background()
	ctx := context.WithValue(c, k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}

func StringsUtils() {

	// strings.Map()
	// strings.Join()
}

func interfaceJson() string {
	return `1`
}

func InterfaceJsonT() {
	var s *interface{}
	jsonStr := interfaceJson()
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
	v, err := json.Marshal(s)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
	var value []byte
	value = []byte{}
	fmt.Println(*s)
	fmt.Println(string(v))
	fmt.Println(len(value))
	fmt.Println(string(value))
}

type Foo2 struct {
}
type Foo struct {
	A string
	B int
	// c int
	D int
	E *string
	F *Foo2
}

func setFoo(foo *Foo, field string, value string) {
	v := reflect.ValueOf(foo).Elem().FieldByName(field)
	if v.IsValid() {
		v.SetString(value)
	}
}

func printFoo() {
	foo := &Foo{
		A: "1",
	}
	s := reflect.ValueOf(foo).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func setFooT() {
	foo := &Foo{
		A: "1",
	}
	fmt.Println(foo)
	setFoo(foo, "A", "2")
	fmt.Println(foo)
}

func ReflectFieldByNameT() {
	printFoo()
	// setFooT()
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

const s = `
# 这是Service1
Services1:
-   Orders:
    -   ID: $save ID1
        SupplierOrderCode: $SupplierOrderCode
    -   ID: $save ID2
        SupplierOrderCode: 111111
# 这是Service2
Services2:
-   Orders:
    -   ID: $save ID1
        SupplierOrderCode: $SupplierOrderCode
    -   ID: $save ID2
        SupplierOrderCode: 111111
`

// func YamlToJSONT() {
// 	fmt.Printf("Input: %s\n", s)
// 	var body interface{}
// 	if err := yaml.Unmarshal([]byte(s), &body); err != nil {
// 		panic(err)
// 	}
//
// 	body = convert(body)
//
// 	if b, err := json.Marshal(body); err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Printf("Output: %s\n", b)
// 	}
// }

type jsonTT struct {
	GRPC       string `json:",omitempty"`
	GRPCUseTLS string `json:",omitempty"`
	CheckID    string `json:",omitempty"`
}

func JsonT() {
	s := `
{
"grpc": "1","gRPC": "2",
"gRPCUseTLS": "1","gRPC_use_tls": "2"
}
`
	var jsonTTT jsonTT
	if err := json.Unmarshal([]byte(s), &jsonTTT); err != nil {
		panic(err)
	} else {
		fmt.Printf("Output: %v\n", jsonTTT)
	}
}

// func JsonToYamlT() {
//	s := `
// {
//    "addr": ":8083",
//    "accesslog":{
//        "filename": "./log/access/ptapollo.log",
//        "request_body": true,
//        "response_body": true
//    },
//    "tracing": {
//        "zipkin": {
//            "http_collector": "http://zipkin.host:9411/api/v1/spans"
//        }
//    },
//	"metrics": {
//		"addr": ":8084"
//	},
//    "consul": {
//        "server": {
//            "address": "http://consul.host:8500",
//            "datacenter": ""
//        },
//        "registration": {
//            "id": "",
//            "tags": ["debug"],
//            "address":"localhost"
//        }
//    },
//    "iswitch_cfg": {
//        "broker":"qconf",
//        "content":{
//          "_qconf_node":"/conf/test/switch"
//        }
//    }
// }
// `
//	var body interface{}
//	if err := json.Unmarshal([]byte(s), &body); err != nil {
//		panic(err)
//	} else {
//		fmt.Printf("Output: %v\n", body)
//	}
//
//	bs, err := yaml.Marshal(body)
//	if err != nil {
//		panic(err)
//	} else {
//		fmt.Printf("Output: %v\n", string(bs))
//	}
// }

func IntConvertT() {
	fmt.Println(int32(2147483647))
}

const (
	MinMaskingLength = 3
	MaskingLength    = 6
)

func Masking(src string, start int, length int, mask string) (des string) {
	if mask == "" {
		mask = "*"
	}
	mask = strings.Repeat(mask, MaskingLength)
	if len(src) <= MinMaskingLength {
		return mask
	}
	des = src[:start] + mask + src[start+length:]
	return
}

func SwitchT() {
	a := []int{1, 2, 3, 4}
	switch {
	case a[0] == 1:
		fmt.Println("a[0]")
		fallthrough
	case a[1] == 2:
		fmt.Println("a[1]")
		fallthrough
	case a[2] == 3:
		fmt.Println("a[2]")
	}

	return
}

func CeilT() {
	v := 1
	// v = v * (90 / 100)
	v = int(math.Ceil(float64(v*90) / 100))
	s := float64(v*90) / 100
	fmt.Println(v)
	fmt.Println(s)
}

func TimeT() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Add(-time.Hour))
	fmt.Println(now.Add(time.Hour))
	now.Add(-time.Hour).Unix()
}

func ErrorT() (err error) {
	err = errors.New("duplicate record")

	defer func() {
		fmt.Printf("record, err<%v>\n", err)
	}()
	if 1 == 1 {
		err := errors.New("1 == 1 error")
		fmt.Printf("1 == 1 err<%v>\n", err)
		return err
	}
	return
}

func timeT() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.AddDate(1, 0, 0))
	fmt.Println(now.AddDate(0, 3, 3))
	fmt.Println(now.AddDate(1, 1, 1))
}

func deferT() {
	// defer deferT()

	/*
		//defer X(a, b, c)
		method = X
		method.args[0] = a
		method.args[1] = b
		method.args[2] = c
	*/

	/*
		//defer X(){ print(a) }
		method = X
	*/

	/*
		method.Invoked(...)
	*/
}

func intPointerT() {
	// var i *int
	// i = append(i, 1)

}

func TimeAddT() {
	fmt.Println(
		time.Date(2020, 1, 19, 23, 59, 59, 0, time.Local).AddDate(0, 0, 31))
}

func GoschedT() {
	go func() {
		fmt.Println(1)
		runtime.Gosched()
		fmt.Println(2)
	}()
	go func() {
		fmt.Println(3)
		runtime.Gosched()
		fmt.Println(4)
	}()
	go func() {
		fmt.Println(5)
		runtime.Gosched()
		fmt.Println(6)
	}()
	/*
		    5
		    1
		    6
		    2
		    3
		    4

			3
		    1
		    4
		    5
		    2
		    6
	*/

	time.Sleep(1 * time.Second)

	// selectT()
	//
	// xx := &a{A: -1, B: -2}
	// fmt.Println(xx)
	// PointerT(xx)
	// PointerT(*xx)
	// fmt.Println(xx)

}

func deferPanicT() {
	defer func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("111")
		}()
		panic(333)
	}()
	// panic(333)
	//
	// defer func() {
	//	fmt.Println(1 / 0)
	// }()
	fmt.Println(4444)
}

func PointerT(xx *a) {
	xx.A = 1
	xx.B = 2
	return
}

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

var closeChan chan byte

type a struct {
	A int
	B int
}

var (
	chanA chan int
	chanB chan int
	chanC chan a
)

func selectT() {
	closeChan = make(chan byte)
	chanA = make(chan int)
	chanB = make(chan int)
	chanC = make(chan a)
	go selectA()
	go selectB()
	go selectC()
	time.Sleep(1 * time.Second)
	// close(closeChan)
	// close(chanA)
	close(chanC)
	time.Sleep(1000 * time.Second)
}

func Close() {
	close(closeChan)
}
func selectA() {
	select {
	case v := <-chanA:
		fmt.Println(v)
	case <-closeChan:
		fmt.Println("close selectA")
		break
	}
}
func selectB() {
	select {
	case v := <-chanB:
		fmt.Println(v)
	case <-closeChan:
		fmt.Println("close selectB")
		break
	}
}
func selectC() {
	select {
	case v := <-chanC:
		fmt.Println(v)
	case <-closeChan:
		fmt.Println("close selectC")
		break
	}
}

type IntError int

const (
	IEOthers IntError = 0
)

func AtomicT() {
	a := atomic.Value{}
	mapX := &map[string]int{
		"1": 1,
		"2": 2,
	}
	a.Store(mapX)
	fmt.Println(a)

	v := a.Load().(*map[string]int)
	fmt.Println(v)
	vv := a.Load().(map[string]int)
	fmt.Println(vv)
}
func UnicodeT() {
	a := "\xFF\xFD"
	b := []byte{0xFF, 0xFD}
	c := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	d := "\xef\xbf\xbd"
	s := "�"
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(string(b))
	fmt.Println(string(b) == s)
	fmt.Println(a == s)
	fmt.Println(d == s)
	fmt.Println([]byte(s))
	fmt.Println([]byte(d))
	fmt.Println(utf8.DecodeRuneInString(s))
	fmt.Println(utf8.DecodeRuneInString(d))

	fmt.Println(strconv.FormatInt(65533, 16))

	var x []byte
	x = make([]byte, 4)
	fmt.Println(utf8.EncodeRune(x, 0xfffd))
	fmt.Println(x)
	fmt.Println(0xfffd == 65533)

	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(unsafe.Sizeof(true))

}

func stringToDateT() {
	layout := "2006-01-02 15:04:05"
	str := "0000-00-00 00:00:00"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	fmt.Println(t.Unix())
}

func timeLocalT() {
	t, _ := time.ParseInLocation("20060102150405", "20200113104349", time.FixedZone("Asia/Shanghai", 8*60*60))
	fmt.Println(t.Unix())
	fmt.Println(time.Unix(t.Unix(), 0))
}

func fileT() {
	filePath := "./file1.txt"
	// fd, err := ioutil.ReadFile(filePath)
	// if err != nil {
	//	fmt.Printf("readFile can't open file: %s\n", filePath)
	//	return
	// }
	// fmt.Println(strings.Split(string(fd), "\n"))
	// fmt.Println([]uint64{8080717372750626869, 8080717373052649480, 8080717372750626869, 8080717372750626869, 8080717373052649480})

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("readFile can't create file: %s\n", filePath)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	n2, err := w.WriteString("asd\n")
	if err != nil {
		fmt.Printf("w.WriteString error: %s\n", err)
		return
	}
	fmt.Println(n2)
	// w.WriteString("哈哈")
	err = w.Flush()
	if err != nil {
		fmt.Printf("w.Flush error: %s\n", err)
		return
	}
	// w.WriteString("456")
	// w.WriteString("789")
	// w.Flush()
	return
}

func NewClientT() {
	c := sdk.NewClient("123", "abc", "http", "sx.api.mengtuiapp.com", 80, 3)

	// GET Request Test
	fmt.Println("GET /v1/zeus/pap_order/8083848945444372559")

	//

	statusCode, resp := c.Get("/v1/zeus/pap_order/8083848945444372559", "pay_method=0&platform=1", nil)

	fmt.Println("Status Code: " + fmt.Sprint(statusCode))
	fmt.Println(resp)

	fmt.Println(strings.Repeat("-", 50))

	// curl -H 'Authorization:mm 123:j7td4yUfkIUma4jRiqv0UmCbWKE=' -v http://sx.api.mengtuiapp.com/v1/zeus/pap_order/8083848945444372559?pay_method=0&platform=1

	// // POST Request Test
	// body := `{"echo":{"int":1,"str":"Hello World","unicode":"你好，世界！","none":null,"boolean":true}}`
	// fmt.Println("GET /x-admin/echo")
	// fmt.Println(body)
	//
	// statusCode, resp = c.Post("/api/v1/do/echo", "", body, nil)
	//
	// fmt.Println("Status Code: " + fmt.Sprint(statusCode))
	// fmt.Println(resp)
}

func copyT() {
	a := []*AA{&AA{x: &BB{x: 1, y: 2}, y: 2}, &AA{x: &BB{x: 1, y: 2}, y: 4}, &AA{x: &BB{x: 1, y: 2}, y: 6}}
	b := append(a[0:0], a...)
	fmt.Printf("*%v: %v\n", a, b)
	fmt.Printf("*%v: %v\n", a[0], b[0])
	fmt.Printf("*%v: %v\n", a[0].x, b[0].x)
	fmt.Printf("*%v: %v\n", a[0].x.x, b[0].x.x)
	a[0].x.x = 11
	a[0].x.y = 11
	fmt.Printf("*%v: %v\n", a[0].x.x, b[0].x.x)

	// fmt.Println(string([]byte{}))
	c := []*CC{&CC{x: BB{x: 1, y: 2}, y: 2}, &CC{x: BB{x: 1, y: 2}, y: 4}, &CC{x: BB{x: 1, y: 2}, y: 6}}
	d := append(c[0:0], c...)
	fmt.Printf("%v: %v\n", c, d)
	fmt.Printf("%v: %v\n", c[0], d[0])
	fmt.Printf("%v: %v\n", c[0].x, d[0].x)
	fmt.Printf("%v: %v\n", c[0].x.x, d[0].x.x)
	c[0].x.x = 11
	c[0].x.y = 11
	fmt.Printf("%v: %v\n", c[0].x.x, d[0].x.x)
}
func timeLocT() {

	// fmt.Println(time.LoadLocation("Asia/Shanghai"))
	// fmt.Println(time.LoadLocation("Asia/Chongqing"))
	// loc, _ := time.LoadLocation("UTC")
	// loc2, _ := time.LoadLocation("Asia/Shanghai")
	// fmt.Println(time.Now().In(loc))
	// fmt.Println(time.Now().In(loc).In(loc2))
	// BeijingLocation := time.FixedZone("Asia/Shanghai", 8*60*60)
	// fmt.Println(time.Now().In(loc).In(BeijingLocation))
	// fmt.Println(time.Now().AddDate(0, 0, -31).Add(+time.Hour * 24 * 3))
	fmt.Println(time.Now())
	fmt.Println(time.Unix(1580973299, 0))
}

// func aseT() error {
// 	req := ""
// 	rb, err := base64.StdEncoding.DecodeString(req)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("rb: %v", rb)
//
// 	apiKey := ""
// 	h := md5.New()
// 	h.Write([]byte(apiKey))
// 	keyStar := []byte(fmt.Sprintf("%x", h.Sum(nil)))
// 	// 用key*对加密串B做AES-256-ECB解密（PKCS7Padding）
// 	_, err = simpleaes.AESDecrypt(rb, keyStar)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

type AA struct {
	Goods []uint64 `json:"goods"`
	x     *BB
	y     int64
}

func (a *AA) add() {
	if a != nil {
		fmt.Println("haha a")
	} else {
		fmt.Println("haha b")
	}
}

type CC struct {
	x BB
	y int64
}

type BB struct {
	x int64
	y int64
}

func nilInCallT() {
	var a *AA
	a.add()
	fmt.Println(a.x)
}

type FuncT struct {
	a int64
}

func (s *FuncT) FuncA(ctx context.Context, b int64) {
	fmt.Printf("FuncA-> a: %v, b: %v\n", s.a, b)
	s.a = 10
	return
}
func (s *FuncT) FuncB(ctx context.Context, b int64) {
	fmt.Printf("FuncB-> a: %v, b: %v\n", s.a, b)
	s.a = 9
	return
}

func funcT() {
	x := &FuncT{}
	funcA := x.FuncA
	funcB := x.FuncB
	x.a = 3

	funcA(context.Background(), 1) // a: 3, b: 1
	funcB(context.Background(), 3) // a: 0, b: 3
	funcA(context.Background(), 1) // a: 10, b: 1
	funcB(context.Background(), 3) // a: 0, b: 3
	x.FuncB(context.Background(), 3)
	funcA(context.Background(), 1) // a: 10, b: 1
	funcB(context.Background(), 3) // a: 0, b: 3

}
func JsonPointT() {
	type A struct {
		B int64 `json:"b"`
		C int64 `json:"c"`
	}

	var s *A
	jsonStr := `{"b": 1, "c": 2}`
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
	fmt.Println(s)

	var b A
	jsonStr = `{"b": 1, "c": 2}`
	err = json.Unmarshal([]byte(jsonStr), &b)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
	fmt.Println(b)

}
func InterfaceT(vMap map[uint]uint) {
	var a interface{}
	if a == nil {
		fmt.Println(1)
	}
}

func MapNilT(vMap map[uint]uint) {
	if vMap == nil {
		fmt.Println("vMap is nil")
	}

	if v, exist := vMap[1]; exist {
		fmt.Println(v)
	} else {
		fmt.Println("dont existed")
	}
}

func GoRoutineT() {
	go func() {
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("1111")
		}()
		fmt.Println("2222")
	}()
	fmt.Println("3333")
	time.Sleep(60 * time.Second)
}

type C int
type DD C

const name DD = 1

func MapNoKeyT() {
	v := map[int64]int64{1: 2, 3: 4}
	fmt.Println(v)
	v[1] += 1
	v[3] += 1
	v[2] += 0
	fmt.Println(v)
}

func deferArgsT() {
	a, b := 1, 2
	defer func(a, b int) {
		fmt.Printf("have args. a: %d, b: %d\n", a, b)
	}(a, b)
	defer func() {
		fmt.Printf("no args. a: %d, b: %d\n", a, b)
	}()

	a = 3
	b = 4
}

func WaitGroupT() {
	var wg *sync.WaitGroup
	wg = &sync.WaitGroup{}
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
	fmt.Println(wg)
	wg.Wait()
	fmt.Println(wg)
}

func f() int {
	return 1
}
func SwitchPrefT() {

	// var x *int32
	// switch x {
	// case nil:
	// 	fmt.Println("111")
	// }
	// >>> 111

	switch x := f(); { // missing switch expression means "true"
	case x < 0:
		fmt.Println(-x)
	default:
		fmt.Println(x)
	}
	switch x := f(); x { // is ok
	case 1:
		fmt.Println(-x)
	case 2, 3:
	default:
		fmt.Println(x)
	}

	// var xx interface{}
	// xx = 1
	// switch i := xx.(type) {
	// case int:
	// 	// i is an int
	// 	fmt.Printf("%T\n", i) // prints "int"
	// case bool:
	// 	// i is a bool
	// 	fmt.Printf("%T\n", i) // prints "bool"
	// 	fallthrough
	// case string:
	// 	fmt.Printf("%T\n", i)
	// 	// What does that type? It should type "string", but if
	// 	// the type was bool and we hit the fallthrough, what would it do then?
	// }

}

func MapPrefT() {
	var a interface{} = 1
	var b interface{} = 1
	fmt.Println(a == b)

	var i = 1
	var x map[int]int = map[int]int{
		i: 2,
		i: 1,
	}
	var xx map[interface{}]int = map[interface{}]int{
		a: 2,
		b: 3,
	}
	fmt.Println(x, xx)

}

func SwitchOrderT() {
	a := 1
	b := 2
	a = 0
	b = 1

	switch {
	case a != 0:
		fmt.Printf("%d\n", a)
	case b != 0:
		fmt.Printf("%d\n", b)
	default:
		return
	}
	fmt.Println("ccc")
}

func deferErrorTDefer() (err error) {
	defer func() {
		err = errors.New("defer")
	}()
	return errors.New("return")
}

func deferErrorTReturn() error {
	var err error
	defer func() {
		err = errors.New("defer")
	}()
	err = errors.New("return")
	return err
}

func deferErrorT() {
	fmt.Println(deferErrorTDefer())  // -> defer
	fmt.Println(deferErrorTReturn()) // -> return
}

func fmtPointPrintT() {
	var deliveryInfoMap map[string]*DeliveryInfo = map[string]*DeliveryInfo{
		"1": &DeliveryInfo{
			OrderDetail:     "1",
			DeliveryNo:      "2",
			DeliveryCompany: "3",
		},
		"2": nil,
	}
	body, err := json.Marshal(deliveryInfoMap)
	if err != nil {
		fmt.Printf("marshal failed. deliveryInfoMap v: %v, err: %v\n", deliveryInfoMap, err)
		return
	}
	fmt.Println(string(body))

	fmt.Printf("deliveryInfoMap v: %v\n", deliveryInfoMap)
	fmt.Printf("deliveryInfoMap +v: %+v\n", deliveryInfoMap)
	fmt.Printf("deliveryInfoMap #v: %#v\n", deliveryInfoMap)
	fmt.Printf("deliveryInfoMap +#v: %+#v\n", deliveryInfoMap)
}

type DeliveryInfo struct {
	OrderDetail     string `json:"orderDetail"`     // 订单Detail号
	DeliveryNo      string `json:"deliveryNo"`      // 物流单号
	DeliveryCompany string `json:"deliveryCompany"` // 物流公司
}
type varNilTFuncClz struct {
	a int8
}

func varNilTFunc() *varNilTFuncClz {
	var v varNilTFuncClz
	return &v
}

func varNilT() {
	fmt.Println(varNilTFunc().a)
}

func tickerT() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("haha1")
			time.Sleep(10 * time.Second)
			fmt.Println("haha")
		}
	}
}

func jsonStringT() {
	var (
		retVal      interface{}
		retJson2Val interface{}
		retValJson  json.RawMessage
		err         error
	)
	retVal = "haha"

	// marshal 结果
	if retVal != nil {
		retValJson, err = json.Marshal(retVal)
		if err != nil {
			fmt.Println(fmt.Sprintf("Run json.Marshal failed with error. retVal: %v, err: %v", retVal, err))
			return
		}
		fmt.Println(fmt.Sprintf("retValJson: %v", string(retValJson)))

		err = json.Unmarshal(retValJson, retJson2Val)
		if err != nil {
			fmt.Println(fmt.Sprintf("Run json.Unmarshal failed with error. retValJson: %v, err: %v", retValJson, err))
			return
		}
		fmt.Println(fmt.Sprintf("retJson2Val: %v", retJson2Val))
	}
	return
}

var iMap = map[string]uint64{
	"1": 1,
}

func poolMapT() {
	productKeplerSkuIdsMap := map[string]int32{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	funcs := make([]func() (interface{}, error), 0, len(productKeplerSkuIdsMap))
	type AX struct {
		ProductId string
	}
	var productId string
	for productId, _ = range productKeplerSkuIdsMap {
		id := productId
		funcs = append(funcs, func() (interface{}, error) {
			productPrice := AX{
				ProductId: id,
			}
			return productPrice, nil
		})
	}

	res, err := PoolMap(funcs)
	if err != nil {
		fmt.Println("PoolMap execute err", err)
		return
	}

	fmt.Println(res)
}

type PoolResult struct {
	idx  int
	Data interface{}
	Err  error
}

func PoolMap(funcs []func() (interface{}, error)) ([]PoolResult, error) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))
	resultChan := make(chan PoolResult, len(funcs))
	for i, fn := range funcs {
		func(idx int, f func() (interface{}, error)) {
			result, err := f()
			resultChan <- PoolResult{
				idx:  idx, // 用 idx 维护是传入函数的结果列表。
				Data: result,
				Err:  err,
			}
			wg.Done()
		}(i, fn)
	}
	retResults := make([]PoolResult, 0, len(funcs))
	wg.Wait()
	close(resultChan)
	for result := range resultChan {
		retResults = append(retResults, result)
	}
	// 将执行结果排序为传入的顺序。
	sort.Slice(retResults, func(i, j int) bool {
		return retResults[i].idx < retResults[j].idx
	})
	return retResults, nil
}

type JAA struct {
	a int
	B time.Time `json:"b"`
}

func (j *JAA) JAAT() {
	fmt.Println("JAAT", j.a)
	j.JAAT2()
}

func (j *JAA) JAAT2() {
	fmt.Println("JAAT2", j.a)
}

type JBB struct {
	JAA
}

func (j *JBB) JAAT() {
	j.JAA.JAAT()
	fmt.Println("JBBT", j.a)
}

func (j *JBB) JAAT2() {
	fmt.Println("JBBT2", j.a)
}

type CommonResponse struct {
	Code    int64           `json:"resultCode"`
	Message string          `json:"resultMessage"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func ExampleUnmarshal() {
	type i struct {
		FullName string `xml:"out_trade_no"`
	}
	var v i

	data := `
		<xml><appid><![CDATA[wx2513074d5de96702]]></appid>
<attach><![CDATA[;5720420194801729624]]></attach>
<bank_type><![CDATA[OTHERS]]></bank_type>
<cash_fee><![CDATA[1]]></cash_fee>
<fee_type><![CDATA[CNY]]></fee_type>
<is_subscribe><![CDATA[N]]></is_subscribe>
<mch_id><![CDATA[1496806362]]></mch_id>
<nonce_str><![CDATA[1e2fc8f62a800391ea83c0c62e3b568c]]></nonce_str>
<openid><![CDATA[oqC3f1BJluSbBp9RVBQsfK_RdGOg]]></openid>
<out_trade_no><![CDATA[t-8026263203931537414-H]]></out_trade_no>
<result_code><![CDATA[SUCCESS]]></result_code>
<return_code><![CDATA[SUCCESS]]></return_code>
<sign><![CDATA[B62AC5444517D7872DD1F5E019E0EFA0]]></sign>
<time_end><![CDATA[20191205210949]]></time_end>
<total_fee>1</total_fee>
<trade_type><![CDATA[MWEB]]></trade_type>
<transaction_id><![CDATA[4200000456201912055517365154]]></transaction_id>
</xml>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%+v", v)
	// Output:
	// XMLName: xml.Name{Space:"", Local:"Person"}
	// Name: "Grace R. Emlin"
	// Phone: "none"
	// Email: [{home gre@example.com} {work gre@work.com}]
	// Groups: [Friends Squash]
	// Address: {Hanga Roa Easter Island}
}

func init() {
	// fmt.Printf("init: %d\n", name)
}

/*

./grpcdebug -data='{"audit_id":"7102968518905217060"}' -addr=127.0.0.1:7401 -method=/ptgold.Service/SucceedBuyMemberV2 {}



./bin/grpcdebug -data='{"order_id_list": [8085968169885286454,8085937970728714294,8085907771991572534,8085877572784668726,8085847373829423158,8085817174756737078,8085786975918932022,8085756776795914294,8085726577823891510,8085696378818314294,8085666179879845942,8085635980857491510,8085605781868691510,8085575582863114294,8085545383958200374,8085515184919068726,8085484985896714294,8085454786907914294,8085424587952668726,8085394388997423158,8085364189958291510,8085333991086932022,8085283675813494838,8085253460416577590,8085223261679435830,8085193062489309238,8085162863450177590,8085132664427823158,8085102465405468726,8085072266483777590,8085042067746635830,8085011868539732022,8084981669685149750,8084951470746681398,8084921271690772534,8084891072550977590,8084860873813835830,8084830674606932022,8084800475752349750,8084770276763549750,8084740077741195318,8084709878601400374,8084669613232668726,8084639414277423158,8084609215406063670,8084579016299823158,8084548817243914294,8084518618305445942,8084488419316645942,8084458220478840886], "pay_method": 0, "platform": 1}' -addr=127.0.0.1:8001 -method=/pay.Service/QueryPapOrderByIdAndPlatform {}
./bin/grpcdebug -data='{"order_id_list": [8085968169885286454, 8085937970728714294], "pay_method": 0, "platform": 1}' -addr=127.0.0.1:8001 -method=/pay.Service/QueryPapOrderByIdAndPlatform {}


./bin/grpcdebug -data='{"order_id":"8093115279084666964","contract_id":"3479407756872564823","buyer_id":36022487,"client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}
./grpcdebug -data='{"order_id":"8093115279084666965","contract_id":"3440644244788789332","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5,"buyer_id":36024584}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}
./grpcdebug -data='{"order_id":"8055290203530641453","contract_id":"3443588436062158893","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5,"buyer_id":36025845}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

test
./grpcdebug -data='{"order_id":"8055290203530641454","contract_id":"3529009359652700197","buyer_id":36034237,"client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}
./grpcdebug -data='{"order_id":"8055290203530641454","contract_id":"3416017396251541558","buyer_id":36022154,"client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

./grpcdebug -data='{"order_id":"8055290203530641454","contract_id":"3412706730803511393","buyer_id":36018997,"client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":1,"order_type":5}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

./bin/grpcdebug -data='{"order_id":"8093115279084666964","contract_id":"3594295742881120284","buyer_id":36022028,"client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

./bin/grpcdebug -data='{"order_id":"8093115279084666968","contract_id":"3594331510110470172","buyer_id":36022028,"client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

./grpcdebug -data='{"order_id":"8085968169885286455","contract_id":"3416017396251541558","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5,"buyer_id":36022154}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}
./grpcdebug -data='{"order_id":"8085968169885286457","contract_id":"3416017396251541558","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5,"buyer_id":36022154}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}


./grpcdebug -data='{"order_id":"8120662902481928290","contract_id":"3412706730803511393","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5,"buyer_id":36018997}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

{
   "order_id": "8093115279084666964",
   "contract_id": "3440644244788789332",
   "client_ip": "39.96.164.183",
   "order_title": "萌推会员自动续费",
   "order_desc":"萌推会员自动续费",
   "platform": 2,
   "order_type": 5,
   "buyer_id": 36024584
}

kubectl cp -n ptpay  ./grpcdebug ptpay-5c8c46f676-rpstl:./


./bin/grpcdebug -data='{"order_id":"8053955264981811294","contract_id":"3442064127841845248","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform":2,"order_type":5,"buyer_id":36022393}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

8030835391066030116,8031273723097104420,8073433748116455460,8074882477562396708,8076331761562337316,8077780224904871972,8079229502865096740,8080678981916966948,8082128305444257828,8083577986846097444,8085027395920314404,8086477153020968996,8087926537113845796,8089375954442420260,8090825444902862884,8092275312280240164,8093724790728212516,8095174268538601508,8118369205883109412

 8073433748116455460, 8079270991896789071, 8080704918318547037, 8080717373052649480, 8083586161527832606, 8095180692115193871, 8115473153827897436


./bin/grpcdebug -data='{"order_id":"3480833734568099928","contract_id":"3472233166979252304","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform": 1,"order_type":5,"buyer_id":36022480}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

./bin/grpcdebug -data='{"order_id":"3479407756872564823","contract_id":"3479407756872564823","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform": 1,"order_type":5,"buyer_id":36022487}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

./bin/grpcdebug -data='{"order_id":"8092522092494372953","contract_id":"3480833734568099928","client_ip":"39.96.164.183","order_title":"萌推会员自动续费","order_desc":"萌推会员自动续费","platform": 1,"order_type":5,"buyer_id":36022488}' -addr=127.0.0.1:8001 -method=/pay.Service/WXPapPayApply {}

106599536,29932779,107799393,107923130,48963208,108234815,29865992

// 订单来源信息;没有进行binding检查，防止参数错误导致下单失败
{
	"ref_page_name": "shijiu-test-ref_page_name-1",
	"ref_page_id": "shijiu-test-ref_page_id-1",
	"ref_key_param": "shijiu-test-ref_key_param-1",
	"source_ref_page_name": "shijiu-test-source_ref_page_name-1",
	"source_ref_page_id": "shijiu-test-source_ref_page_id-1",
	"source_ref_key_param": "shijiu-test-source_ref_key_param-1",
    "source_ref_pos_id": "shijiu-test-source_ref_pos_id-1",
	"activity_id": 662,
	"extra": {
		"extra_key_1": "extra_value_1",
		"extra_key_2": "extra_value_2",
		"extra_key_3": "extra_value_3",
	},
}
*/
