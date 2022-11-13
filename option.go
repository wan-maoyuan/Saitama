package saitama

import (
	"context"
	"os"
	"time"

	"github.com/wanmaoyuan/saitama/transport"
)

type Option func(o *options)

type options struct {
	ctx         context.Context
	name        string             // server name
	version     string             // server version
	stopTimeout time.Duration      // server stop timeout
	signals     []os.Signal        // server stop singals
	servers     []transport.Server // any server list

	beforeStart func(context.Context) error // before server start run func
	beforeStop  func(context.Context) error // before server stop run func
	afterStart  func(context.Context) error // after server start run func
	afterStop   func(context.Context) error // after server stop run func
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

func WithStopTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.stopTimeout = timeout
	}
}

func WithSignals(singals ...os.Signal) Option {
	return func(o *options) {
		o.signals = singals
	}
}

func WithServers(servers ...transport.Server) Option {
	return func(o *options) {
		o.servers = servers
	}
}

func WithBeforeStart(beforeStart func(context.Context) error) Option {
	return func(o *options) {
		o.beforeStart = beforeStart
	}
}

func WithBeforeStop(beforeStop func(context.Context) error) Option {
	return func(o *options) {
		o.beforeStop = beforeStop
	}
}

func WithAfterStart(afterStart func(context.Context) error) Option {
	return func(o *options) {
		o.afterStart = afterStart
	}
}

func WithAfterStop(afterStop func(context.Context) error) Option {
	return func(o *options) {
		o.afterStop = afterStop
	}
}
