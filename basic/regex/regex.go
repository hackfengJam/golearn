package main

import (
	"fmt"
	"regexp"
)

//const text = "My email is hackfun@163.com"
const text = `My email is hackfun@163.com
email1 abc@163.com
email2 kkk@163.com
email2 ddd@163.com.cn
`

func main() {
	//re := regexp.MustCompile("hackfun@163\\.com")
	//re := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9.]+\\.[a-zA-Z0-9]+")
	re := regexp.MustCompile("([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\\.[a-zA-Z0-9.]+)")
	//re := regexp.MustCompile(`hackfun@163\.com`)
	//match := re.FindString(text)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
	//fmt.Println(match)
}
