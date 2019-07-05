package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Printf("%p \n", &stu)
		m[stu.Name] = &stu
	}

	fmt.Println(m)
	fmt.Println(m["zhou"].Age, m["li"].Age, m["wang"].Age, )
}
func pase_student_cor() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i := range stus {
		stu := stus[i]
		m[stu.Name] = &stu
	}

	fmt.Println(m)
	fmt.Println(m["zhou"].Age, m["li"].Age, m["wang"].Age, )
}

func pase_student_cor2() {
	m := make(map[string]*student)
	stus := []*student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = stu
		//m[stu.Name] = &stu
	}

	fmt.Println(m)
	fmt.Println(m["zhou"].Age, m["li"].Age, m["wang"].Age, )
}

func main() {
	// map[zhou:0xc42000a060 li:0xc42000a060 wang:0xc42000a060]
	// 22 22 22
	pase_student()
	// map[zhou:0xc42000a060 li:0xc42000a060 wang:0xc42000a060]
	// 24 23 22
	pase_student_cor()
	// map[zhou:0xc42000a060 li:0xc42000a060 wang:0xc42000a060]
	// 24 23 22
	pase_student_cor2()
}
