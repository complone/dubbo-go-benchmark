package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
)

var (
	userProvider = &UserProvider{}
)

func main() {

	config.SetConsumerService(userProvider)
	hessian.RegisterPOJO(&User{})

	path := "/Users/windwheel/Documents/gitrepo/dubbo-go-benchmark/3.0/pixiu/client/dubbogo.yml"
	err := config.Load(config.WithPath(path))

	if err != nil {
		panic(err)
	}
	user, err := userProvider.GetUserByName(context.TODO(), "chengxingyuan")
	if user != nil {
		fmt.Println("当前对象为: {}", user)
	}

}
