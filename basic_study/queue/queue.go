package queue

// 1
//type Queue []interface{}
//func (q *Queue) Push(v interface{}) {
//	*q = slice_append(*q, v)
//}
//
//func (q *Queue) Pop() interface{} {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head
//}
//
//func (q *Queue) IsEmpty() bool {
//	return len(*q) == 0
//}

// 2
//type Queue []interface{}
//
//func (q *Queue) Push(v int) {
//	*q = slice_append(*q, v)
//}
//
//func (q *Queue) Pop() int {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head.(int)
//}
//
//func (q *Queue) IsEmpty() bool {
//	return len(*q) == 0
//}

// 3
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
