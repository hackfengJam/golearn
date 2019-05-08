package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type resource struct {
	url    string
	target string
	start  int
	end    int
}

func ruleResource() []resource {
	var res [] resource
	r1 := resource{ // 首页
		url:    "http://localhost:8888/",
		target: "",
		start:  0,
		end:    0,
	}
	r2 := resource{ // 列表页
		url:    "http://localhost:8888/list/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    21,
	}
	r3 := resource{ // 详情页
		url:    "http://localhost:8888/movie/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    12924,
	}
	res = append(append(append(res, r1), r2), r3)
	return res
}
func buildUrl(res [] resource) []string {
	var list []string
	for _, resItem := range res {

		if len(resItem.target) == 0 {
			list = append(list, resItem.url)
		} else {
			for i := resItem.start; i <= resItem.end; i++ {
				urlStr := strings.Replace(resItem.url, resItem.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}
func makeLog(current, refer, ua string) string {
	u := url.Values{}
	u.Set("time", "1")
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)

	paramsStr := u.Encode()
	logTemplate := ""
	log := strings.Replace(logTemplate, "{$paramsStr}", paramsStr, -1)
	log = strings.Replace(log, "{${ua}}", ua, -1)
	return log
}

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Intn(max-min) + min
}

func main() {
	total := flag.Int("total", 100, "how many rows by created")
	filePath := flag.String("filePath", "/User/Hack/nginx/logs/dig.log", "log file path")
	flag.Parse()
	var uaList = []string{}
	//fmt.Println(*total, *filePath)
	// 需要构造出真是的网站url集合
	res := ruleResource()
	list := buildUrl(res)
	//fmt.Println(list)

	//按照要求，生成 $total 行日志内容，源自上面的这个集合
	logStr := ""
	for i := 0; i <= *total; i++ {
		currentUrl := list [randInt(0, len(list)-1)]
		referUrl := list [randInt(0, len(list)-1)]
		ua := uaList[randInt(0, len(uaList)-1)]

		logStr = logStr + makeLog(currentUrl, referUrl, ua) + "\n"
		//ioutil.WriteFile(*filePath, []byte(logStr), 0644)
	}
	fd, _ := os.OpenFile(*filePath, os.O_RDWR|os.O_APPEND, 0644)
	fd.Write([]byte(logStr))
	fd.Close()

	fmt.Println("done.")

}
