package global

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	counterValue uint64
	ipInfo       sync.Map
)

// Info ...
type Info struct {
	IP        string `json:"ip"`
	BucketLen int    `json:"len"`
	BucketCap int    `json:"cap"`
}

// Add counter加1
func Add() {
	atomic.AddUint64(&counterValue, 1)
}

// Num 取得目前數量
func Num() uint64 {
	return atomic.LoadUint64(&counterValue)
}

// StoreIPInfo ...
func StoreIPInfo(ip string, info *Info) {
	ipInfo.Store(ip, info)
}

// GetIPInfo ...
func GetIPInfo(ip string) (info *Info, err error) {
	data, ok := ipInfo.Load(ip)
	if !ok {
		// set info
		return nil, errors.New("get ip info failed")
	}
	return data.(*Info), nil
}

// GetAllIPInfos ...
func GetAllIPInfos() interface{} {
	ipInfos := make([]*Info, 0)
	iter := func(key interface{}, value interface{}) bool {
		// ip := key.(string)
		info := value.(*Info)
		// ipInfos[ip] = info
		ipInfos = append(ipInfos, info)
		return true
	}
	ipInfo.Range(iter)
	return ipInfos
}
