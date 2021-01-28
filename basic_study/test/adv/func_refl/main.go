package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

func runTaskFunc(taskFunc *reflect.Value, message *TaskMessage) (*ResultMessage, error) {

	// check number of arguments
	numArgs := taskFunc.Type().NumIn()
	messageNumArgs := len(message.Args)
	if numArgs != messageNumArgs+1 {
		return nil, fmt.Errorf("number of task arguments %d does not match number of message arguments %d", numArgs, messageNumArgs)
	}
	// TODO
	ctx := context.Background()
	// construct arguments
	in := make([]reflect.Value, messageNumArgs+1)
	in[0] = reflect.ValueOf(ctx)
	for i, arg := range message.Args {
		index := i + 1

		origType := taskFunc.Type().In(index).Kind()
		msgType := reflect.TypeOf(arg).Kind()
		// special case - convert float64 to int if applicable
		// this is due to json limitation where all numbers are converted to float64
		if msgType == reflect.Float64 {
			switch origType {
			case reflect.Int:
				arg = int(arg.(float64))
			case reflect.Int8:
				arg = int8(arg.(float64))
			case reflect.Int16:
				arg = int16(arg.(float64))
			case reflect.Int32:
				arg = int32(arg.(float64))
			case reflect.Int64:
				arg = int64(arg.(float64))
			case reflect.Uint:
				arg = uint(arg.(float64))
			case reflect.Uint8:
				arg = uint8(arg.(float64))
			case reflect.Uint16:
				arg = uint16(arg.(float64))
			case reflect.Uint32:
				arg = uint32(arg.(float64))
			case reflect.Uint64:
				arg = uint64(arg.(float64))
			default:
				// do nothing
			}
		}

		in[index] = reflect.ValueOf(arg)
	}

	// call method
	res := taskFunc.Call(in)
	if len(res) == 0 {
		return nil, nil
	}

	return getReflectionResultMessage(&res[0]), nil
}

type Res struct {
	C float64 `json:"c"`
	D int64   `json:"d"`
	E string  `json:"e"`
	Args
}
type Args struct {
	C float64 `json:"c"`
	D int64   `json:"d"`
	E string  `json:"e"`
}

// func Add(ctx context.Context, a float64, b int64, c string, args Args) *Res {
//     return &Res{
//         C: a,
//         D: b,
//         E: c,
//         Args: Args{
//             C: args.C,
//             D: args.D,
//             E: args.E,
//         },
//     }
// }

func Add(ctx context.Context, a float64, b int64, c string) *Res {
	return &Res{
		C: a,
		D: b,
		E: c,
	}
}

func main() {
	// s := `{"id": "1", "task": "2", "args": [1,2,"a", {"c":1,"d":2,"e":"abc"}]}`
	s := `{"id": "1", "task": "2", "args": [1,2,"a"]}`
	msg := taskMessagePool.Get().(*TaskMessage)
	err := json.Unmarshal([]byte(s), msg)
	if err != nil {
		panic(err)
	}
	// use reflection to execute function ptr
	taskFunc := reflect.ValueOf(Add)
	res, err := runTaskFunc(&taskFunc, msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", res.Result)
	v, ok := res.Result.(*reflect.Value)
	if ok {
		k := v.Kind()
		p := (v).Pointer()

		fmt.Printf("ok, v: %v\n", v)
		fmt.Printf("ok, vk: %v\n", k)
		fmt.Printf("ok, *v: %v\n", *v)
		fmt.Printf("ok, *vk: %v\n", (*v).Kind())
		T(*v)
		fmt.Printf("ok, vp: %v\n", p)
		fmt.Printf("ok, *vp: %v\n", (*v).Pointer())
		fmt.Printf("\n")
		fmt.Printf("ok, ptr: %v\n", p)
		fmt.Printf("ok, Interface: %v\n", v.Interface())
		fmt.Printf("ok, *vInterface: %v\n", (*(v).Interface().(*Res)))
	}

	// v, ok := res.Result.(*Res)
	// if ok {
	// 	fmt.Println(v)
	// }

}
func T(v reflect.Value) {
	fmt.Printf("T(v reflect.Value), *v: %v\n", v)
	fmt.Printf("T(v reflect.Value), *vk: %v\n", v.Kind())
	fmt.Printf("T(v reflect.Value), *v: %v\n", v)
	fmt.Printf("T(v reflect.Value), *vk: %v\n", *((*Res)(unsafe.Pointer(v.Pointer()))))
}
