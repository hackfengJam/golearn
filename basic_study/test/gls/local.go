package gls

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var gls struct {
	m map[int64]map[interface{}]interface{}
	sync.Mutex
}

func init() {
	gls.m = make(map[int64]map[interface{}]interface{})
}

// https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch3-asm/ch3-08-goroutine-id.html
func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}

func getMap() map[interface{}]interface{} {
	gls.Lock()
	defer gls.Unlock()

	goid := GetGoid()
	if m, _ := gls.m[goid]; m != nil {
		return m
	}

	m := make(map[interface{}]interface{})
	gls.m[goid] = m
	return m

}

func Get(key interface{}) interface{} {
	return getMap()[key]
}

func Put(key interface{}, v interface{}) {
	getMap()[key] = v
}

func Delete(key interface{}) {
	delete(getMap(), key)
}

func Clean() {
	gls.Lock()
	defer gls.Unlock()

	delete(gls.m, GetGoid())
}
