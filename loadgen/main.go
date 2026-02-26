package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	target := "http://localhost:8080/work"

	rps := 10
	maxRPS := 150

	for rps <= maxRPS {
		fmt.Println("Current RPS:", rps)

		var wg sync.WaitGroup

		for i := 0; i < rps; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				client.Get(target)
			}()
		}

		wg.Wait()

		time.Sleep(1 * time.Second)
		rps += 5
	}
}