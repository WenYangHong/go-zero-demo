package tool

import (
	"sync"
	"time"
)

const (
	workerIdBits     = uint(5)
	datacenterIdBits = uint(5)
	sequenceBits     = uint(12)

	maxWorkerId     = -1 ^ (-1 << workerIdBits)
	maxDatacenterId = -1 ^ (-1 << datacenterIdBits)

	workerIdShift      = sequenceBits
	datacenterIdShift  = sequenceBits + workerIdBits
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits
	sequenceMask       = -1 ^ (-1 << sequenceBits)

	twepoch = int64(1672531200000) // 2023-01-01
)

type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerId      int64
	datacenterId  int64
	sequence      int64
}

var defaultNode *Snowflake
var once sync.Once

// Init 初始化（推荐在 main 或 svc 中调用）
func Init(workerId, datacenterId int64) {
	once.Do(func() {
		defaultNode = New(workerId, datacenterId)
	})
}

// New 创建一个独立实例（用于特殊场景）
func New(workerId, datacenterId int64) *Snowflake {
	if workerId > maxWorkerId || workerId < 0 {
		panic("workerId out of range")
	}
	if datacenterId > maxDatacenterId || datacenterId < 0 {
		panic("datacenterId out of range")
	}

	return &Snowflake{
		workerId:     workerId,
		datacenterId: datacenterId,
	}
}

// NextId 获取ID（全局默认实例）
func NextId() int64 {
	if defaultNode == nil {
		panic("snowflake not initialized")
	}
	return defaultNode.next()
}

// NextId 获取ID（实例方法）
func (s *Snowflake) NextId() int64 {
	return s.next()
}

// 内部实现
func (s *Snowflake) next() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1e6

	if now == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}

	// 时钟回拨保护（简单版）
	if now < s.lastTimestamp {
		panic("clock moved backwards")
	}

	s.lastTimestamp = now

	return ((now - twepoch) << timestampLeftShift) |
		(s.datacenterId << datacenterIdShift) |
		(s.workerId << workerIdShift) |
		s.sequence
}
