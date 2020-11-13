package ratelimiter

import "time"

// Option ...
type Option interface {
	apply(*RateLimiter)
}

type optFunc func(*RateLimiter)

func (f optFunc) apply(r *RateLimiter) {
	f(r)
}

// LimitTime 設置限制時間
func LimitTime(d time.Duration) Option {
	return optFunc(func(r *RateLimiter) {
		r.limitDuartion = d
	})
}

// LimitNum 設置限制數量
func LimitNum(num int) Option {
	return optFunc(func(r *RateLimiter) {
		r.limitNum = num
	})
}
