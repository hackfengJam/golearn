package main

import (
	"fmt"
	"unsafe"
)

func mapT()  {
	data := make(map[int64]string,2)

	data[1] = "aaa"
	data[2] = "bbb"
	data[3] = "ccc"
	data[12312312] = "dddd"

	fmt.Println(data)

	for _,v := range data{
		oldV := v
		ln := len(v)-1
		v= v[0:ln]+"0"+v[ln+1:]
		fmt.Printf("%s ", v)
		fmt.Printf("%s ", oldV)
	}

	fmt.Println()
	fmt.Println(data)

}

func arrT()  {
	data := "abcdefgh"
	fmt.Println(data[:len(data)])

}

func sprintFT()  {
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

func mapAdT()  {
	data := map[string]string{"a":"1","b":"2","c":"3",}
	fmt.Println(data)
	set2Map(data, "a", "2")
	set2Map(data, "d", "4")
	fmt.Println(data)
}
func mulZero()  {

	str:= "%0"+ fmt.Sprintf("%d", 3)+"s"
	fmt.Println(fmt.Sprintf(str, "0"))
}

func sliceT(){
	a := []int{1,2,3}
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
	fmt.Printf("%p %p", unsafe.Pointer(&a[0]),b)

}

func IntT()  {
	//i := 10
	//i = 2 / 3
	a:= []int{1,2,3,4,5}
	//segCount := len(a) /
	fmt.Println(a[:(len(a) / 3 )* 3])

}

func main() {
	//mapT()
	//arrT()
	//sprintFT()
	//mapAdT()
	//sliceT()
	IntT()
	//v := []int{1,2,3}
	//fmt.Println(v[:0])
}
