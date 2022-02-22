/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

type UserProviderProxy struct {
	userProvider *UserProvider
}

//func (u *UserProvider) GetUser(ctx context.Context, userID *Request) (*User, error) {
//
//	json, _ := json.Marshal(&userID)
//	if  json!= nil{
//
//	}
//	fmt.Println("发起调用请求: %s",json)
//	return u.GetUser(ctx, userID)
//	//return &User{
//	//	ID: "12345",
//	//	Name: "Hello" + userID,
//	//	Age: 21,
//	//}, nil
//}

var (
	userProvider = &UserProvider{}
)

// need to setup environment variable "DUBBO_GO_CONFIG_PATH" to "conf/dubbogo.yml" before run
func main() {

	config.SetConsumerService(userProvider)

	path := "/Users/windwheel/Documents/gitrepo/dubbo-go-benchmark/3.0/dubbo/client/dubbogo.yml"
	err := config.Load(config.WithPath(path))

	if err != nil {
		panic(err)
	}
	reqUser := &Request{
		Name: "12345",
	}

	user, err := userProvider.GetUser(context.TODO(), reqUser)

	if err != nil {
		panic(err)
	}
	logger.Infof("response result: %v", user)

	//config.SetProviderService(&UserProviderProxy{x
	//	userProvider: userProvider,
	//})

}
