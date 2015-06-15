package main

import (
	"fmt"
	"sync/atomic"
)

// 分钟级别统计，每个domain一个
type PeerMinuStat struct {
	fetchNum      int64 // 一分钟之内，所有代理的抓取数量
	errorFetchNum int64 // 一分钟之内，所有代理的抓取失败数量
}

type PeerStat struct {
	ncount  int64                    // 总的抓取数量
	nerror  int64                    // 总的抓取失败数量
	hostMap map[string]*PeerMinuStat //每个host的每分钟的抓取状况
}

// 创建一个新的统计
func NewPeerStat() *PeerStat {
	peerStat := &PeerStat{
		ncount:  0,
		nerror:  0,
		hostMap: make(map[string]*PeerMinuStat),
	}

	return peerStat
}

func (p *PeerStat) add(host string) {
	atomic.AddInt64(&p.ncount, 1)
	_, ok := p.hostMap[host]
	if !ok {
		p.hostMap[host] = &PeerMinuStat{}
	}
	atomic.AddInt64(&p.hostMap[host].fetchNum, 1)
}

func main() {
	peerStat := NewPeerStat()
	peerStat.add("dianping.com")
	fmt.Println(peerStat)
	fmt.Println(peerStat.hostMap["dianping.com"])

	var test int64 = 67
	atomic.StoreInt64(&test, 10)
	fmt.Println(test)

}
