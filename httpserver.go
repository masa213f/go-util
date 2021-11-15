package util

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

const HeaderRequestID = "X-Request-Id"

type wrapResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *wrapResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

type HTTPServer struct {
	log    logr.Logger
	server *http.Server
}

func NewHTTPServer(log logr.Logger, addr string, handler http.Handler) *HTTPServer {
	requestLogger := func(rw http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(HeaderRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
			r.Header.Set(HeaderRequestID, requestID)
		}
		wrap := &wrapResponseWriter{rw, http.StatusOK}

		start := time.Now()
		handler.ServeHTTP(wrap, r)
		end := time.Now()

		msec := (end.Sub(start)).Microseconds()
		log.Info("request",
			"request_id", requestID,
			"elapsed_time", msec,
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"status_code", wrap.statusCode,
			"referer", r.Referer(),
			"user_agent", r.UserAgent(),
		)
	}

	return &HTTPServer{
		log:    log,
		server: &http.Server{Addr: addr, Handler: http.HandlerFunc(requestLogger)},
	}
}

func (h *HTTPServer) Serve(ctx context.Context) error {
	errCh := make(chan error)
	go func() {
		h.log.Info("listen", "addr", h.server.Addr)
		if err := h.server.ListenAndServe(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
	case err := <-errCh:
		return fmt.Errorf("failed to listen: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown: %v", err)
	}
	h.log.Info("shutdown", "addr", h.server.Addr)
	return nil
}
