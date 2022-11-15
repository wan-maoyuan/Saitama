package saitama

import (
	"context"
	"fmt"
	"syscall"
	"testing"
	"time"
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

func TestNew(t *testing.T) {
	sa := New(
		WithName("test"),
		WithVersion("v1.0.0"),
		WithStopTimeout(time.Second*5),
		WithSignals(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT),
		WithServers(&TestServer1{}),
		WithBeforeStart(beforeStartFunc),
		WithBeforeStop(beforeStopFunc),
		WithAfterStart(afterStartFunc),
		WithAfterStop(afterStopFunc),
	)

	go sa.Run()
	time.Sleep(time.Second)
	sa.Stop()
}
