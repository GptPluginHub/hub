package middleware

import (
	"net/http"
	"time"

	"k8s.io/klog/v2"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

const (
	ReadyzURLPath  = "/readyz"
	HealthzURLPath = "/healthz"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func (r *ResponseWriter) WriteHeader(code int) {
	r.StatusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func LogRequestAndResponse(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ws := NewResponseWriter(w)

		// The Grpc-Metadata-X-Real-IP header is added to get the Real IP and port in the grpc middleware
		r.Header.Set("Grpc-Metadata-X-Real-IP", r.RemoteAddr)
		handler.ServeHTTP(ws, r)
		if DetermineTheSpecialPath(r.URL.Path) {
			klog.Infof("[kangaroo] %s |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n",
				time.Now().Format("2006-01-02 - 15:04:05"),
				StatusCodeColor(ws.StatusCode), ws.StatusCode, ResetColor(),
				time.Since(start).String(),
				r.RemoteAddr,
				MethodColor(r.Method), r.Method, ResetColor(),
				r.URL.Path,
			)
		}
	})
}

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func StatusCodeColor(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

// MethodColor is the ANSI color for appropriately logging http method to a terminal.
func MethodColor(method string) string {
	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}

// ResetColor resets all escape attributes.
func ResetColor() string {
	return reset
}

// DetermineTheSpecialPath check if path is a value inside const.
func DetermineTheSpecialPath(path string) bool {
	switch path {
	case ReadyzURLPath, HealthzURLPath:
		return false
	}
	return true
}
