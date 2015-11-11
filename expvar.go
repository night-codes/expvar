package expvar

import (
	"expvar"
	"net/http"
	"runtime"
	"time"
)

var startTime = time.Now().UTC()

func goroutines() interface{} {
	return runtime.NumGoroutine()
}

// uptime is an expvar.Func compliant wrapper for uptime info.
func uptime() interface{} {
	uptime := time.Since(startTime)
	return int64(uptime)
}

// Start expvar export
func Start(host string) {
	expvar.Publish("Goroutines", expvar.Func(goroutines))
	expvar.Publish("Uptime", expvar.Func(uptime))
	go http.ListenAndServe(host, nil) //debug port
}

