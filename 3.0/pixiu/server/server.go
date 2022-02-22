package main

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var survivalTimeout = int(3e9)

var (
	userProvider = &UserProvider{}
)

// they are necessary:
// 		export CONF_PROVIDER_FILE_PATH="xxx"
// 		export APP_LOG_CONF_FILE="xxx"
func main() {

	cache = newUserDB()
	t1, _ := time.Parse(
		time.RFC3339,
		"2021-08-01T10:08:41+00:00")

	cache.Add(&User{ID: "0001", Code: 1, Name: "tc", Age: 18, Time: t1})
	cache.Add(&User{ID: "0002", Code: 2, Name: "ic", Age: 88, Time: t1})
	config.SetProviderService(userProvider)
	hessian.RegisterPOJO(&User{})

	// ------for hessian2------

	path := "/Users/windwheel/Documents/gitrepo/dubbo-go-benchmark/3.0/pixiu/server/dubbogo.yml"
	if err := config.Load(config.WithPath(path)); err != nil {
		panic(err)
	}
	initSignal()
}

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}
