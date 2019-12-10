package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AkConfigArgs struct {
	Title string `json:"title"` // 标题
	Name  string `json:"name"`  // name 唯一
	Value string `json:"value"` // 值
}

type Permissions struct {
	Permissions int `json:"permissions"` // 权限 1：读 2：写
}
type Role struct {
	Id string `json:"id"`
}

type ReqT struct {
	AkConfigArgs AkConfigArgs  `json:"ak_config"`
	Role         Role          `json:"role"`
	Permissions  []Permissions `json:"permissions"`
}

func main() {
	pwds := `{
		"test-22252": "123qaz"
	}`
	ignoreValue := false

	var pwdMap map[string]string
	err := json.Unmarshal([]byte(pwds), &pwdMap)
	if err != nil {
		panic(err)
	}

	// data
	var list []AkConfigArgs
	for k, v := range pwdMap {
		list = append(list, AkConfigArgs{
			Title: "",
			Name:  k,
			Value: v,
		})
	}

	// req
	url := "http://192.168.18.59:30488/x-admin/general/ak_config"
	headers := map[string]string{}
	client := &http.Client{}
	var (
		req  *http.Request
		resp *http.Response
	)

	// post
	for _, akConfig := range list {
		// construct req
		if ignoreValue {
			akConfig.Value = ""
		}
		reqT := ReqT{
			AkConfigArgs: akConfig,
			Permissions: []Permissions{
				{Permissions: 1}, // 读
			},
		}
		b, _ := json.Marshal(reqT)
		req, err = http.NewRequest(
			http.MethodPost,
			url, bytes.NewBuffer(b))
		if err != nil {
			panic(err)
		}

		// add header
		for k, v := range headers {
			req.Header.Add(k, v)
		}

		// req
		resp, err = client.Do(req)
		if err != nil {
			panic(err)
		}

		// resp
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// log
		fmt.Printf("k-> %s, status:%s, resp: %v \n", akConfig.Name, resp.Status, string(data))
	}

	// close
	if resp != nil {
		_ = resp.Body.Close()
	}
}
