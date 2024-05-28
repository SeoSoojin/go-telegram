package webhook

type RouteOptions func(*routeOptions)

type routeOptions struct {
	callbacks map[string]Command
	aliases   []string
}

func WithCallbacks(callbacks map[string]Command) RouteOptions {
	return func(o *routeOptions) {
		o.callbacks = callbacks
	}
}

func WithCallback(name string, callback Command) RouteOptions {
	return func(o *routeOptions) {
		o.callbacks[name] = callback
	}
}

func WithAliases(aliases ...string) RouteOptions {
	return func(o *routeOptions) {
		o.aliases = aliases
	}
}

func newRouteOptions(opts ...RouteOptions) *routeOptions {
	options := &routeOptions{
		callbacks: make(map[string]Command),
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
