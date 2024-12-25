package locks

import "runtime"

type PoolOptions struct {
	Size int
}

func DefaultOptions() *PoolOptions {
	return &PoolOptions{
		Size: runtime.GOMAXPROCS(0) * 10,
	}
}

func (p *PoolOptions) Modify(opts ...PoolOption) {
	for _, o := range opts {
		o.apply(p)
	}
}

func (p *PoolOptions) Sanitize() {
	p.Size = max(p.Size, 10)
}

type PoolOption interface {
	apply(*PoolOptions)
}

type poolOptionFn func(*PoolOptions)

func (modify poolOptionFn) apply(opts *PoolOptions) {
	modify(opts)
}

func WithSize(size int) PoolOption {
	return poolOptionFn(func(po *PoolOptions) {
		po.Size = size
	})
}
