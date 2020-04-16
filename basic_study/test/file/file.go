package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("/Users/vesper/Data/WorkSpaces/golang/GOPATH/src/pt-cluster/src/ptapp.cn/tool/internal/grpcdebug")))
}
