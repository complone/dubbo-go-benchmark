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
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	tester2 "github.com/dubbogo/tools/pkg/tester"
	"os"
	"strconv"
)

var (
	userProvider = &UserProvider{}
)

const (
	Tps      = "TPS"
	Parallel = "PARALLEL"
	// Duration should be a string representing a time,
	// like "1h", "30m", etc.
	Duration = "DURATION"
	FuncName = "FUNC_NAME"

	// Supported FuncNames
	Fibonacci = "FIBONACCI"
	Sleep     = "SLEEP"
	CLIENT    = "CLIENT"

	// FuncName == "FIBONACCI"
	FibonacciN         = "FIBONACCI_N"
	FibonacciWorkerNum = "FIBONACCI_WORKER_NUM"

	// FuncName == "SLEEP"
	SleepDuration = "SLEEP_DURATION"
)

func init() {
	hessian.RegisterPOJO(&User{})
	hessian.RegisterPOJO(&Request{})
}

// need to setup environment variable "DUBBO_GO_CONFIG_PATH" to "conf/dubbogo.yml" before run
func main() {
	config.SetConsumerService(userProvider)
	if err := config.Load(config.WithPath("/Users/windwheel/Documents/gitrepo/dubbo-go-benchmark/3.0/dubbo/stress/dubbogo.yml")); err != nil {
		panic(err)
	}

	var (
		tps, parallel      int
		duration, funcName string
		err                error
	)

	ctx := context.TODO()
	if tps, err = strconv.Atoi(os.Getenv(Tps)); err != nil {
		panic(fmt.Errorf("env %s is required: %w", Tps, err))
	}
	logger.Infof("TPS is set to %d.", tps)
	if parallel, err = strconv.Atoi(os.Getenv(Parallel)); err != nil {
		panic(fmt.Errorf("env %s is required: %w", Parallel, err))
	}
	logger.Infof("Parallel is set to %d.", parallel)
	if duration = os.Getenv(Duration); duration == "" {
		panic(fmt.Errorf("%s is required", Duration))
	}
	if funcName = os.Getenv(FuncName); funcName == "" {
		panic(fmt.Errorf("%s is required", FuncName))
	}

	reqUser := &Request{
		Name: "yiya yiya",
	}

	doInvoke := func(uid int) {
		switch funcName {

		case CLIENT:
			user, err := userProvider.GetUser(ctx, reqUser)
			if err != nil {
				fmt.Sprintf("client resonse: %s", user)
			}

		default:
			panic(fmt.Sprintf("%s is an unsupported function", funcName))
		}
	}

	tester := tester2.NewStressTester()
	tester.
		SetTPS(tps).
		SetDuration(duration).
		SetTestFn(doInvoke).
		SetUserNum(parallel).
		Run()

	fmt.Printf("Sent request num: %d", tester.GetTransactionNum())
	fmt.Printf("TPS: %.2f\n", tester.GetTPS())
	fmt.Printf("RT: %.2fs\n", tester.GetAverageRTSeconds())

	//ctx := context.Background()
	//tpsNum, _ := strconv.Atoi(os.Getenv("tps"))
	//parallel, _ := strconv.Atoi(os.Getenv("parallel"))
	//payloadLen, _ := strconv.Atoi(os.Getenv("payload"))
	//req  := "laurence" + string(make([]byte, payloadLen))
	//stressTest.NewStressTestConfigBuilder().SetTPS(tpsNum).SetDuration("1h").SetParallel(parallel).Build().Start(func() {
	//	if _, err := userProvider.GetUser(ctx, &Request{
	//		Name: req,
	//	}); err != nil{
	//		panic(err)
	//	}
	//})
}
