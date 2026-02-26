package main

func ComputeSlope(samples []Sample) float64 {
	if len(samples) < 2 {
		return 0
	}

	first := samples[0]
	last := samples[len(samples)-1]

	deltaLatency := last.Latency - first.Latency
	deltaTime := last.Timestamp.Sub(first.Timestamp).Seconds()

	if deltaTime == 0 {
		return 0
	}

	return deltaLatency / deltaTime
}