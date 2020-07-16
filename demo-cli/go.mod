module laracom/demo-cli

go 1.14

replace laracom/demo-service => /Users/wangqi/go/src/microservice/demo/src/laracom/demo-service

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/go-micro v1.18.0
	laracom/demo-service v0.0.0-00010101000000-000000000000
)
