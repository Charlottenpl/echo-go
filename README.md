# echo-go
### 构建
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
### 运行
sudo nohup ./echo-go > nohup_echo-go.log 2>&1 &
ps -ef|grep main
