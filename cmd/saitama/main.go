package main

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/wanmaoyuan/saitama"
)

func beforeStartFunc(ctx context.Context) error {
	fmt.Println("before start function run")
	return nil
}

func beforeStopFunc(ctx context.Context) error {
	fmt.Println("before stop function run")
	return nil
}

func afterStartFunc(ctx context.Context) error {
	fmt.Println("after start function run")
	return nil
}

func afterStopFunc(ctx context.Context) error {
	fmt.Println("after stop function run")
	return nil
}

type TestServer1 struct{}

func (s *TestServer1) Start(ctx context.Context) error {

	return nil
}

func (s *TestServer1) Stop(ctx context.Context) error {

	return nil
}

func main() {
	sa := saitama.New(
		saitama.WithName("test"),
		saitama.WithVersion("v1.0.0"),
		saitama.WithStopTimeout(time.Second*5),
		saitama.WithSignals(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT),
		saitama.WithServers(&TestServer1{}),
		saitama.WithBeforeStart(beforeStartFunc),
		saitama.WithBeforeStop(beforeStopFunc),
		saitama.WithAfterStart(afterStartFunc),
		saitama.WithAfterStop(afterStopFunc),
	)

	sa.Run()
}
