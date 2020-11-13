package ratelimiter

import (
	"fmt"
	"sync"
	"time"

	"github.com/miles990/ratelimiter-example/global"
)

// RateLimiter ...
type RateLimiter struct {
	limitDuartion time.Duration
	limitNum      int

	throttles map[string]chan *global.Info
	mu        sync.Mutex
}

// NewRateLimiter ...
func NewRateLimiter(opts ...Option) *RateLimiter {
	// default
	r := &RateLimiter{
		limitDuartion: time.Second,
		limitNum:      5,
		throttles:     make(map[string]chan *global.Info, 5),
	}

	for _, opt := range opts {
		opt.apply(r)
	}
	return r
}

// Check 確認是否被限流，如果被限流回傳 error
func (r *RateLimiter) Check(ip string) (bufferLen int, bufferCap int, err error) {
	global.Add()
	// 如果 channel buffer 已滿不做事
	_, ok := r.throttles[ip]
	if !ok {
		r.mu.Lock()
		r.throttles[ip] = make(chan *global.Info, r.limitNum)
		r.mu.Unlock()
	}
	bufferLen = len(r.throttles[ip])
	bufferCap = cap(r.throttles[ip])

	// 判斷 channel 是否已滿
	if bufferLen == bufferCap {
		fmt.Println(fmt.Sprintf("bucket full [%v]", ip))
		return bufferLen, bufferCap, ErrOverLimit
	}

	// 如果 channel 未滿，buffer + 1
	info := &global.Info{
		IP:        ip,
		BucketLen: bufferLen + 1,
		BucketCap: bufferCap,
	}
	r.throttles[ip] <- info
	global.StoreIPInfo(ip, info)
	fmt.Println(fmt.Sprintf("bucket add [%v] %+v", ip, info))

	// 釋放 channel buffer
	time.AfterFunc(r.limitDuartion, func() {
		data := <-r.throttles[ip]
		changeIPInfo := &global.Info{
			IP:        data.IP,
			BucketLen: len(r.throttles[data.IP]),
			BucketCap: len(r.throttles[data.IP]),
		}
		global.StoreIPInfo(data.IP, changeIPInfo)
		fmt.Println(fmt.Sprintf("bucket release [%v] %+v", ip, changeIPInfo))
	})
	return bufferLen, bufferCap, nil
}
