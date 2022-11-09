package saitama

import (
	"context"
	"os"
)

type Option func(o *options)

type options struct {
	name    string      // server name
	version string      // server version
	signals []os.Signal // server stop singals

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

func WithSignals(singals ...os.Signal) Option {
	return func(o *options) {
		o.signals = singals
	}
}
