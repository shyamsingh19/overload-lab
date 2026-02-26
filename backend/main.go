package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var inFlight int64

const (
	baseLatency     = 50 * time.Millisecond
	perRequestDelay = 5 * time.Millisecond
	maxCapacity     = 100 // simulated capacity
)

func handler(w http.ResponseWriter, r *http.Request) {
	current := atomic.AddInt64(&inFlight, 1)
	defer atomic.AddInt64(&inFlight, -1)

	// Simulate nonlinear delay near saturation
	utilization := float64(current) / float64(maxCapacity)

	delay := baseLatency + time.Duration(current)*perRequestDelay

	// Add nonlinear explosion near saturation
	if utilization > 0.8 {
		extra := time.Duration((utilization-0.8)*500) * time.Millisecond
		delay += extra
	}

	time.Sleep(delay)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/work", handler)
	fmt.Println("Backend running on :8080")
	http.ListenAndServe(":8080", nil)
}