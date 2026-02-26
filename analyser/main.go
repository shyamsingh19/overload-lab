package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	window := NewSlidingWindow(5)
	client := &http.Client{Timeout: 2 * time.Second}

	ticker := time.NewTicker(2 * time.Second)

	for range ticker.C {
		start := time.Now()
		_, err := client.Get("http://localhost:8080/work")
		if err != nil {
			fmt.Println("Probe failed:", err)
			continue
		}
		latency := time.Since(start).Seconds() * 1000 // ms

		sample := Sample{
			Timestamp: time.Now(),
			Latency:   latency,
		}

		window.Add(sample)

		slope := ComputeSlope(window.Samples())

		fmt.Printf("Latency: %.2f ms | Slope: %.4f\n", latency, slope)

		if slope > 20 { // threshold tuning later
			fmt.Println("Instability detected")
		}
	}
}