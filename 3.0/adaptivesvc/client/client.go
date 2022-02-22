package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	tester2 "github.com/dubbogo/tools/pkg/tester"
	"os"
	"strconv"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
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

	// FuncName == "FIBONACCI"
	FibonacciN         = "FIBONACCI_N"
	FibonacciWorkerNum = "FIBONACCI_WORKER_NUM"

	// FuncName == "SLEEP"
	SleepDuration = "SLEEP_DURATION"
)

var (
	provider = &FibProvider{}
)

func main() {

	config.SetConsumerService(provider)
	path := "/Users/windwheel/Documents/gitrepo/dubbo-go-benchmark/3.0/adaptivesvc/client/dubbogo.yml"
	if err := config.Load(config.WithPath(path)); err != nil {
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

	doInvoke := func(uid int) {
		switch funcName {
		case Fibonacci:
			if result, err := fibonacci(ctx, provider); err != nil {
				panic(err)
			} else {
				fmt.Printf("%s result: %d\n", Fibonacci, result)
			}
		case Sleep:
			sleep(ctx, provider)
			fmt.Printf("sleep task was finished")
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
	//
	doInvoke(1)

	//// TODO(justxuewei): remove after test
	//if err := config.Load(); err != nil {
	//	panic(err)
	//}
}

func fibonacci(ctx context.Context, provider *FibProvider) (result int64, err error) {
	var (
		n, workNum int
	)
	if n, err = strconv.Atoi(os.Getenv(FibonacciN)); err != nil {
		panic(err)
	}
	if workNum, err = strconv.Atoi(os.Getenv(FibonacciWorkerNum)); err != nil {
		panic(err)
	}

	result, err = provider.Fibonacci(ctx, int64(n), int64(workNum))
	return
}

func sleep(ctx context.Context, provider *FibProvider) {
	var (
		duration int
		err      error
	)
	if duration, err = strconv.Atoi(os.Getenv(SleepDuration)); err != nil {
		panic(err)
	}
	_, _ = provider.Sleep(ctx, int64(duration))
}
