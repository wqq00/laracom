module laracom/user-cli

go 1.14

require (
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	laracom/user-service v0.0.0-00010101000000-000000000000
)

replace laracom/user-service => /Users/wangqi/go/src/microservice/demo/src/laracom/user-service

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
