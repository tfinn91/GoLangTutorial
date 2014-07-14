package main

import (
	"github.com/etsy/statsd/examples/go"
)

func main() {
	// Record a start time
	//t1 := time.Now()

	// Create a new StatsD connection
	host := "localhost"
	port := 8125
	client := statsd.New(host, port)



	// Increment a stat counter
	client.Increment("stat.metric1")
	client.Increment("stat.metric2")

	// Decrement a stat counter
	client.Decrement("stat.metric1")



	// Record an end time
	//t2 := time.Now()

	// Submit timing information
	//duration := int64(t2.Sub(t1) / time.Millisecond)
	//client.Timing("stat.timer", duration)
}
