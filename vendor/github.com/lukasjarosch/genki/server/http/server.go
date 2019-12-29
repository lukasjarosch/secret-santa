package http

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/etherlabsio/healthcheck"

	"github.com/lukasjarosch/genki/logger"
	"github.com/lukasjarosch/genki/server/http/middleware"
)

type HttpServer struct {
	opts       Options
	httpServer *http.Server
}

func NewServer(opts ...Option) Server {
	options := newOptions(opts...)

	return &HttpServer{
		opts: options,
	}
}

// Handle registers the given endpoint pattern and handler to the HttpServer's http multiplexer.
// If the HttpServer has middleware enabled, then the enabled middleware is also added.
// The call order is the inverse registration call order.
// The 'Metadata' middleware is always the first middleware. It cannot be disabled.
func (srv *HttpServer) Handle(endpoint string, handler http.Handler) {
	if srv.opts.LoggingMiddlewareEnabled {
		handler = middleware.Logging(handler, srv.opts.LoggingSkipEndpoints...)
		logger.Debugf("HTTP HttpServer '%s': logging middleware enabled for endpoint '%s'", srv.opts.Name, endpoint)
	}
	if srv.opts.PrometheusMiddlewareEnabled {
		handler = middleware.Prometheus(handler, endpoint)
		logger.Debugf("HTTP HttpServer '%s': prometheus middleware enabled for endpoint '%s'", srv.opts.Name, endpoint)
	}
	handler = middleware.Metadata(handler)
	logger.Debugf("HTTP HttpServer '%s': metadata middleware enabled for endpoint '%s'", srv.opts.Name, endpoint)
	srv.opts.Multiplexer.Handle(endpoint, handler)
}

// HandlerFunc attaches the given handlerFunc to the HttpServer's multiplexer.
func (srv *HttpServer) HandleFunc(endpoint string, handlerFunc http.HandlerFunc) {
	if srv.opts.LoggingMiddlewareEnabled {
		handlerFunc = middleware.LoggingFunc(handlerFunc)
		logger.Debugf("HTTP HttpServer '%s': logging middleware enabled for endpoint '%s'", srv.opts.Name, endpoint)
	}
	srv.opts.Multiplexer.HandleFunc(endpoint, handlerFunc)
}

// ListenAndServe the HttpServer in a separate goroutine.
// Will block until the context is done.
func (srv *HttpServer) ListenAndServe(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// add health endpoint
	srv.Handle(srv.opts.HealthEndpoint, healthcheck.Handler())
	logger.Infof("registered /health HTTP HttpServer '%s'", srv.opts.Name)

	srv.httpServer = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", srv.opts.Port),
		Handler: srv.opts.Multiplexer,
	}

	// serve
	go func() {
		logger.Infof("HTTP HttpServer '%s' running on port %s", srv.opts.Name, srv.opts.Port)
		if err := srv.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("HTTP HttpServer crashed: %s", err.Error())
			return
		}
	}()

	<-ctx.Done()
	srv.shutdown()
}

// shutdown gracefully terminates the HttpServer with the configured grace period timeout.
func (srv *HttpServer) shutdown() {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), srv.opts.ShutdownGracePeriod)
	defer cancel()
	if err := srv.httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Warnf("HTTP '%s' graceful shutdown timed-out: %s", srv.opts.Name, err.Error())
	} else {
		logger.Infof("HTTP HttpServer '%s' shut-down gracefully", srv.opts.Name)
	}
}

type Option func(*Options)
