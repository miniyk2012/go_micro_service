package privateConfig

type privateConfig struct {
	a string
	b string
	c int
	d int
}

type MyOption interface {
	apply(*privateConfig)
}

type myOptionFunc struct {
	f func(*privateConfig)
}

func (f *myOptionFunc) apply(c *privateConfig) {
	f.f(c)
}

func newOptionFunc(f func(*privateConfig)) MyOption {
	return &myOptionFunc{
		f: f,
	}
}

func WithB(b string) MyOption {
	return newOptionFunc(func(cf *privateConfig) {
		cf.b = b
	})

}

func WithC(c int) MyOption {
	return newOptionFunc(func(cf *privateConfig) {
		cf.c = c
	})
}

func WithA(a string) MyOption {
	return newOptionFunc(func(cf *privateConfig) {
		cf.a = a
	})
}

func NewConfigWithOptions(opts ...MyOption) *privateConfig {
	cf := &privateConfig{
		d: 20,
	}
	for _, opt := range opts {
		opt.apply(cf)
	}
	return cf
}

type FuncOption func(*privateConfig)

func WithD(d int) FuncOption {
	return func(cf *privateConfig) {
		cf.d = d
	}
}

// NewConfigWithFuncOptions 简洁写法, 不定义接口, 定义一个函数FuncOption
func NewConfigWithFuncOptions(opts ...FuncOption) *privateConfig {
	cf := &privateConfig{
		c: 10,
	}
	for _, opt := range opts {
		opt(cf)
	}
	return cf
}
