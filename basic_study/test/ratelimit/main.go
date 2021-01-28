package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type RateLimit interface {
	Init(N, M int64) error
	Consume() (bool, error)
}

type IRateLimit struct {
	N int64 // 每秒 n 个令牌
	// M      int64 // init m 个令牌
	Token  int64     // 桶里 token 数
	LastTs time.Time // 最后一次计算 token 数 时间点

	Lock sync.Locker
}

func New(N, M int64) RateLimit {
	return &IRateLimit{
		N: N,
		// M:      0,
		Token:  M,
		LastTs: time.Now(),
		Lock:   &sync.Mutex{},
	}
}

// 并发不安全
func (rl *IRateLimit) Init(N, M int64) error {
	rl.N = N
	rl.LastTs = time.Now()
	rl.Token = M

	return nil
}

func (rl *IRateLimit) diffTimeSecond(now time.Time) (sec int64) {
	return int64(now.Second() - rl.LastTs.Second())
}

// 并发不安全
func (rl *IRateLimit) refreshToken() {
	now := time.Now()
	// FIXME 溢出暂不考虑
	rl.incrToken(rl.diffTimeSecond(now) * rl.N)
	rl.LastTs = now
	return
}

func (rl *IRateLimit) consumeToken() (ok bool) {
	// lock
	rl.Lock.Lock()
	// unlock
	defer rl.Lock.Unlock()

	if rl.getToken() > 0 {
		rl.incrToken(-1)
		return true
	}

	rl.refreshToken()

	if rl.getToken() > 0 {
		rl.incrToken(-1)
		return true
	}

	return false
}

func (rl *IRateLimit) getToken() (Token int64) {
	return atomic.LoadInt64(&rl.Token)
}

func (rl *IRateLimit) incrToken(Token int64) {
	atomic.AddInt64(&rl.Token, Token)
	return
}

func (rl *IRateLimit) Consume() (bool, error) {
	return rl.consumeToken(), nil
}

func T() {

	// 每秒 n 个令牌

	// init m 个令牌

	// m < n

	//
}

func main() {

}
