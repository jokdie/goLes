package group

import (
	"context"
	"sync"
)

type Group struct {
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	firstError error
	once       sync.Once
}

func NewGroup(parentCtx context.Context) *Group {
	ctx, cancel := context.WithCancel(parentCtx)

	return &Group{
		wg:     sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (g *Group) Go(f func(ctx context.Context) error) {
	g.wg.Go(func() {
		select {
		case <-g.ctx.Done():
			return
		default:
			if err := f(g.ctx); err != nil {
				g.once.Do(func() {
					g.firstError = err
					g.cancel()
				})
			}

			return
		}
	})
}

func (g *Group) Wait() error {
	g.wg.Wait()

	return g.firstError
}
