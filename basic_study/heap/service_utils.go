package heap

import (
	"fmt"
)

type Goods struct {
	GoodsId     int
	OriginPrice int
}

type GoodsHeap struct {
	goodsList []*Goods
	size      int
}

// 实现heap.Interface需要的5个方法
func (m GoodsHeap) Len() int {
	return len(m.goodsList)
}

func (m GoodsHeap) Less(i, j int) bool {
	return m.goodsList[i].OriginPrice < m.goodsList[j].OriginPrice
}

func (m GoodsHeap) Swap(i, j int) {
	m.goodsList[i], m.goodsList[j] = m.goodsList[j], m.goodsList[i]
}

// Push和Pop都要用指针接收者, 因为要在函数内修改slice
func (m *GoodsHeap) Push(x interface{}) {
	m.goodsList = append(m.goodsList, x.(*Goods))
}

func (m *GoodsHeap) Pop() interface{} {
	old := m.goodsList
	n := len(old)
	x := old[n-1]
	m.goodsList = old[0 : n-1]
	return x
}

func (m GoodsHeap) String() string {
	str := "["
	for i, v := range m.goodsList {
		str += fmt.Sprintf("GoodsId: %v OriginPrice: %v", v.GoodsId, v.OriginPrice)
		if i != m.Len()-1 {
			str += ", "
		}
	}
	str += "]"
	str += fmt.Sprintf(" %v", m.Len())
	return str
}
