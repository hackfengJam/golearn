package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golearn/basic_study/test/gls"
	"golearn/basic_study/test/sdk"
	"gopkg.in/yaml.v2"
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
func mulChar() {
	fmt.Println(strings.Repeat("*", 3))
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

func NewClientT() {
	c := sdk.NewClient("ak-7QKrWXH8QZf3O8Tf", "40ZGi4fZsNhB9WaYE3fYdI9cBsievlL3", "http", "127.0.0.1", 8081, 3)

	// GET Request Test
	fmt.Println("GET /api/v1/do/ping")

	statusCode, resp := c.Get("/x-admin/ping", "", nil)

	fmt.Println("Status Code: " + fmt.Sprint(statusCode))
	fmt.Println(resp)

	fmt.Println(strings.Repeat("-", 50))

	// POST Request Test
	body := `{"echo":{"int":1,"str":"Hello World","unicode":"你好，世界！","none":null,"boolean":true}}`
	fmt.Println("GET /x-admin/echo")
	fmt.Println(body)

	statusCode, resp = c.Post("/api/v1/do/echo", "", body, nil)

	fmt.Println("Status Code: " + fmt.Sprint(statusCode))
	fmt.Println(resp)
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

	//strings.Map()
	//strings.Join()
}

func interfaceJson() string {
	return `1`
}

func InterfaceJsonT() {
	var s *interface{}
	jsonStr := interfaceJson()
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	v, err := json.Marshal(s)
	if err != nil {
		//fmt.Println(err)
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
	//c int
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
	//setFooT()
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

func YamlToJSONT() {
	fmt.Printf("Input: %s\n", s)
	var body interface{}
	if err := yaml.Unmarshal([]byte(s), &body); err != nil {
		panic(err)
	}

	body = convert(body)

	if b, err := json.Marshal(body); err != nil {
		panic(err)
	} else {
		fmt.Printf("Output: %s\n", b)
	}
}

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

//func JsonToYamlT() {
//	s := `
//{
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
//}
//`
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
//}

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
	//v = v * (90 / 100)
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
	//defer deferT()

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
	var i *int
	i = append(i, 1)

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

	//var a []*int
	//for i := range a {
	//	fmt.Println("1", i)
	//}
	//fmt.Println(1)

	//JsonMapInterfaceT()
	//JsonMapStructT()
	//regexpT()
	//bitOrT()
	//shiftRightLogicalT()
	//gRpcDebugT()

	//Authorization4SignT()
	//NewClientT()
	//SortSliceT()
	//DateFormatT()
	//MapIfT()
	//NewT()
	//ContextT()
	//StringsUtils()
	//InterfaceJsonT()
	//ReflectFieldByNameT()
	//JsonT()
	//mulChar()
	//CeilT()
	//YamlToJSONT()
	//JsonToYamlT()
	//SwitchT()
	//TimeT()
	//_ = ErrorT()
	//timeT()

	deferT()

	return
}
