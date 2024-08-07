package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var meter = otel.Meter("my-service-meter")

func main() {
	var err error
	itemsCounter, err := meter.Int64UpDownCounter(
		"items.counter",
		metric.WithDescription("Number of items."),
		metric.WithUnit("{item}"),
	)
	if err != nil {
		panic(err)
	}

	_ = func() {
		// code that adds an item to the collection
		itemsCounter.Add(context.Background(), 1)
	}

	_ = func() {
		// code that removes an item from the collection
		itemsCounter.Add(context.Background(), -1)
	}
}
