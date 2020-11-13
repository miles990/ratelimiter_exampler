# ratelimiter example
利用 golang channel 的特性做簡易限流，預設同ip訪問位置，每分鐘接受訪問為 60 個 request

# Build From source 或直接執行 /bin 內執行檔
```
mkdir -p bin/windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows/ratelimiter_exampler.exe main.go

mkdir -p bin/linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/ratelimiter_exampler main.go

mkdir -p bin/osx
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/osx/ratelimiter_exampler main.go
```

# curl
```
curl [ip]:[port]
```

# 未被限流回傳
```
{"count":17,"info":[{"ip":"::1","len":17,"cap":60}],"time":"2020-11-13T13:22:18.2020516+08:00"}
```

# 被限流回傳
```
{"count":61,"err":"Error"}
```

# Usage
```
  -help
        Show this help
  -num int
        limit num (default 60)
  -port int
        http server listen port (default 80)
  -t int
        limit time (second) (default 60)
```