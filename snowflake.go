package snowflake

import (
	"math"
	"sync"
	"time"
)

// 雪花结构体
type SnwoFlake struct {
	DataCenterId int64
	MachineId    int64
	MachineBits  int64
	// 起始时间，时间戳都是对这个时间的差
	StartEpoch int64
	Mutex      sync.Mutex
}

var lastStamp int64 = -1
var sequence int64 = 0

func NewSnowFlake(dataCenterId int64, MachineId int64, machineBits int64, startEpoch int64) *SnwoFlake {
	if machineBits <= 0 || machineBits >= 10 {
		panic("Parameter error!")
	}
	return &SnwoFlake{
		MachineBits: machineBits,
		StartEpoch:  startEpoch,
	}
}

// 核心方法，生成下一个 ID
func (sf *SnwoFlake) NextId() int64 {
	sf.Mutex.Lock()
	defer sf.Mutex.Unlock()
	curStamp := time.Now().Unix()
	// 防止时针反转
	if curStamp < lastStamp {
		panic("Clock moved backwards! Refuse to generate id")
	}
	// 如果在上次生成在同一秒
	if curStamp == lastStamp {
		sequence = (sequence + 1) & int64(math.Pow(2, float64(12))-1)
		if sequence == 0 {
			curStamp = waitNextMill()
		}
	} else {
		sequence = 0
	}
	lastStamp = curStamp
	// 通过移位操作构造 ID
	return (curStamp-sf.StartEpoch)<<22 |
		sf.DataCenterId<<(12+sf.MachineBits) |
		sf.MachineId<<12 |
		sequence
}

// 空循环直到进入下一毫秒
func waitNextMill() int64 {
	mill := time.Now().Unix()
	for mill <= lastStamp {
		mill = time.Now().Unix()
	}
	return mill
}
