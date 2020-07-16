module github.com/wqq00/laracom/user-cli

go 1.14

replace github.com/wqq00/laracom/user-service => /Users/wangqi/go/src/microservice/demo/src/laracom/user-service

require (
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/wqq00/laracom/user-service v0.0.0-00010101000000-000000000000
)
