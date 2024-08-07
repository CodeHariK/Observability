package main

import (
	"context"
	"fmt"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var meter = otel.Meter("my-service-meter")

func main() {
	// This is just a sample of memory stats to record from the Memstats
	heapAlloc, err := meter.Int64ObservableUpDownCounter("heapAllocs")
	if err != nil {
		fmt.Println("failed to register updown counter for heapAllocs")
		panic(err)
	}
	gcCount, err := meter.Int64ObservableCounter("gcCount")
	if err != nil {
		fmt.Println("failed to register counter for gcCount")
		panic(err)
	}

	_, err = meter.RegisterCallback(
		func(_ context.Context, o metric.Observer) error {
			memStats := &runtime.MemStats{}

			// This call does work
			runtime.ReadMemStats(memStats)

			o.ObserveInt64(heapAlloc, int64(memStats.HeapAlloc))
			o.ObserveInt64(gcCount, int64(memStats.NumGC))

			return nil
		},
		heapAlloc,
		gcCount,
	)
	if err != nil {
		fmt.Println("Failed to register callback")
		panic(err)
	}
}
