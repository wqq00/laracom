package test

import (
	"fmt"
	api "laracom/demo-service/api"
)

const (
	appName = "Demo Service"
)

func main()  {
	fmt.Printf("Starting %v\n", appName)
	var port = "8000"
	api.StartWebServer(port)
}