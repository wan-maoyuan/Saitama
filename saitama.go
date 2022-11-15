package saitama

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Saitama struct {
	ctx    context.Context
	cancel func()
	opts   options
}

func New(opts ...Option) *Saitama {
	o := options{
		ctx:         context.Background(),
		stopTimeout: time.Second * 5,
		signals:     []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(o.ctx)

	return &Saitama{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
	}
}

func (sa *Saitama) Run() error {
	if err := sa.beforeStartFunc(sa.ctx); err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(sa.ctx)
	wg := sync.WaitGroup{}

	// start all server
	for _, svr := range sa.opts.servers {
		svr := svr
		eg.Go(func() error {
			<-ctx.Done()
			stopCtx, cancel := context.WithTimeout(context.Background(), sa.opts.stopTimeout)
			defer cancel()

			return svr.Stop(stopCtx)
		})
		wg.Add(1)

		eg.Go(func() error {
			wg.Done()
			return svr.Start(ctx)
		})
	}
	wg.Wait()

	if err := sa.afterStartFunc(sa.ctx); err != nil {
		return err
	}

	// monitor system signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, sa.opts.signals...)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case <-c:
			return sa.Stop()
		}
	})

	if err := eg.Wait(); err != nil && errors.Is(err, context.Canceled) {
		return err
	}

	return sa.afterStopFunc(context.Background())
}

func (sa *Saitama) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), sa.opts.stopTimeout)
	defer cancel()

	if err := sa.beforeStopFunc(ctx); err != nil {
		return err
	}

	return nil
}

func (sa *Saitama) beforeStartFunc(ctx context.Context) error {
	if sa.opts.beforeStart != nil {
		if err := sa.opts.beforeStart(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (sa *Saitama) afterStartFunc(ctx context.Context) error {
	if sa.opts.afterStart != nil {
		if err := sa.opts.afterStart(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (sa *Saitama) beforeStopFunc(ctx context.Context) error {
	if sa.opts.beforeStop != nil {
		if err := sa.opts.beforeStop(ctx); err != nil {
			if sa.cancel != nil {
				sa.cancel()
			}

			return err
		}
	}

	return nil
}

func (sa *Saitama) afterStopFunc(ctx context.Context) error {
	if sa.opts.afterStop != nil {
		if err := sa.opts.afterStop(ctx); err != nil {
			return err
		}
	}

	return nil
}
