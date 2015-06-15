package main

import (
	"container/heap"
	"fmt"
)

type Peer struct {
	load  int
	index int
}

// 继承了堆结构，实现了堆的方法
type peerQueue []*Peer

// 堆成员函数
func (pq peerQueue) Len() int { return int(len(pq)) }

// 比较堆两个成员的负载大小
func (pq peerQueue) Less(i, j int) bool {
	return pq[i].load < pq[j].load
}

// 交换两个node，来确保堆的结构
func (pq peerQueue) Swap(i, j int) {
	fmt.Println("before swap i:", i, " j:", j, " pq[i].index:", pq[i].index, " pq[j].index:", pq[j].index)
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
	fmt.Println("after swap i:", i, " j:", j, " pq[i].index:", pq[i].index, " pq[j].index:", pq[j].index)
}

// 把x放到队尾
func (pq *peerQueue) Push(x interface{}) {
	n := len(*pq)
	peer := x.(*Peer)
	peer.index = n
	*pq = append(*pq, peer)
}

// 把队尾的元素返回
func (pq *peerQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func main() {
	peers := make(peerQueue, 0)
	for i := 0; i < 10; i++ {
		peer := &Peer{
			load: 0,
		}
		heap.Push(&peers, peer)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("index: ", peers[i].index, "  load:", peers[i].load)
	}

	fmt.Println("\n\n")
	peers[0].load = 4
	heap.Fix(&peers, 0)

	for i := 0; i < 10; i++ {
		fmt.Println("index: ", peers[i].index, "  load:", peers[i].load)
	}

	peers[6].load = 0
	heap.Fix(&peers, 6)

	fmt.Println("\n\n")

	for i := 0; i < len(peers); i++ {
		fmt.Println("index: ", peers[i].index, "  load:", peers[i].load)
	}
}
