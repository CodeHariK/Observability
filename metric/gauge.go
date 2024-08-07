package main

import (
	"context"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var meter = otel.Meter("my-service-meter")

func main() {
	speedGauge, err := meter.Int64Gauge(
		"cpu.fan.speed",
		metric.WithDescription("Speed of CPU fan"),
		metric.WithUnit("RPM"),
	)
	if err != nil {
		panic(err)
	}

	getCPUFanSpeed := func() int64 {
		// Generates a random fan speed for demonstration purpose.
		// In real world applications, replace this to get the actual fan speed.
		return int64(1500 + rand.Intn(1000))
	}

	fanSpeedSubscription := make(chan int64, 1)
	go func() {
		defer close(fanSpeedSubscription)

		for idx := 0; idx < 5; idx++ {
			// Synchronous gauges are used when the measurement cycle is
			// synchronous to an external change.
			// Simulate that external cycle here.
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fanSpeed := getCPUFanSpeed()
			fanSpeedSubscription <- fanSpeed
		}
	}()

	ctx := context.Background()
	for fanSpeed := range fanSpeedSubscription {
		speedGauge.Record(ctx, fanSpeed)
	}
}
