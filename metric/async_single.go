package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var meter = otel.Meter("my-service-meter")

func main() {
	_, err := meter.Int64ObservableGauge(
		"DiskUsage",
		metric.WithUnit("By"),
		metric.WithInt64Callback(func(_ context.Context, obsrv metric.Int64Observer) error {
			// Do the real work here to get the real disk usage. For example,
			//
			//   usage, err := GetDiskUsage(diskID)
			//   if err != nil {
			//   	if retryable(err) {
			//   		// Retry the usage measurement.
			//   	} else {
			//   		return err
			//   	}
			//   }
			//
			// For demonstration purpose, a static value is used here.
			usage := 75000
			obsrv.Observe(int64(usage), metric.WithAttributes(attribute.Int("disk.id", 3)))
			return nil
		}),
	)
	if err != nil {
		fmt.Println("failed to register instrument")
		panic(err)
	}
}
