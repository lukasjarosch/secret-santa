package http

import (
	"net/http"
	"time"

	"github.com/spf13/pflag"
)

const DefaultPort = "8080"
const DefaultGracePeriod = 3 * time.Second
const DefaultName = "default"
const DefaultLoggingMiddlewareEnabled = true
const DefaultPrometheusMiddlewareEnabled = true
const DefaultHealthEndpoint = "/health"

type Options struct {
	Port                        string
	Name                        string
	HealthEndpoint              string
	ShutdownGracePeriod         time.Duration
	LoggingMiddlewareEnabled    bool
	PrometheusMiddlewareEnabled bool
	LoggingSkipEndpoints        []string
	Multiplexer            *http.ServeMux
}

func Port(addr string) Option {
	return func(opts *Options) {
		opts.Port = addr
	}
}

func ShutdownGracePeriod(duration time.Duration) Option {
	return func(opts *Options) {
		opts.ShutdownGracePeriod = duration
	}
}

func Name(name string) Option {
	return func(opts *Options) {
		opts.Name = name
	}
}

func DisableLoggingMiddleware() Option {
	return func(opts *Options) {
		opts.LoggingMiddlewareEnabled = false
	}
}

func LoggingSkipEndpoints(skip ...string) Option {
	return func(opts *Options) {
		opts.LoggingSkipEndpoints = append(opts.LoggingSkipEndpoints, skip...)
	}
}

func HealthEndpoint(endpoint string) Option {
	return func(opts *Options) {
		opts.HealthEndpoint = endpoint
	}
}

func Multiplexer(mux *http.ServeMux) Option {
	return func(opts *Options) {
		opts.Multiplexer = mux
	}
}


func newOptions(opts ...Option) Options {
	opt := Options{
		Port:                        DefaultPort,
		ShutdownGracePeriod:         DefaultGracePeriod,
		Name:                        DefaultName,
		LoggingMiddlewareEnabled:    DefaultLoggingMiddlewareEnabled,
		PrometheusMiddlewareEnabled: DefaultPrometheusMiddlewareEnabled,
		HealthEndpoint:              DefaultHealthEndpoint,
		LoggingSkipEndpoints:        []string{},
		Multiplexer:            http.NewServeMux(),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Flags is a convenience function to quickly add the HTTP HttpServer options as CLI flags
// Implements the cli.FlagProvider type
func Flags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("http-HttpServer", pflag.ContinueOnError)

	fs.String(
		PortConfigKey,
		DefaultPort,
		"the port on which the HTTP HttpServer is listening on",
	)
	fs.Duration(
		GracePeriodConfigKey,
		DefaultGracePeriod,
		"grace period after which the HttpServer shutdown is terminated",
	)

	return fs
}
