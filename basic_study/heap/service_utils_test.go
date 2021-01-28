package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestGoodsHeap(t *testing.T) {
	//ast := assert.New(t)

	var goodsHeap *GoodsHeap
	goodsHeap = &GoodsHeap{
		size: 2,
	}
	goodsList := []*Goods{
		{
			GoodsId:     1,
			OriginPrice: 100,
		},
		{
			GoodsId:     2,
			OriginPrice: 0,
		},
		{
			GoodsId:     3,
			OriginPrice: 99,
		},
		{
			GoodsId:     4,
			OriginPrice: 1,
		},
	}

	// sorted by OriginPrice
	expectedGoodsList := []*Goods{
		{
			GoodsId:     2,
			OriginPrice: 0,
		},
		{
			GoodsId:     4,
			OriginPrice: 1,
		},
		{
			GoodsId:     3,
			OriginPrice: 99,
		},
		{
			GoodsId:     1,
			OriginPrice: 100,
		},
	}
	heap.Init(goodsHeap)
	for _, v := range goodsList {
		heap.Push(goodsHeap, v)
		fmt.Println(goodsHeap)
	}

	for i := range expectedGoodsList {
		//ast.Equal(heap.Pop(goodsHeap), expectedGoodsList[i])
		i = i
	}

}
