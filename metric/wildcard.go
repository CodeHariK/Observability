package main

import (
	"fmt"

	"go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	// Create a view that sets unit to milliseconds for any instrument with a
	// name suffix of ".ms".
	view := metric.NewView(
		metric.Instrument{Name: "*.ms"},
		metric.Stream{Unit: "ms"},
	)

	// The created view can then be registered with the OpenTelemetry metric
	// SDK using the WithView option.
	_ = metric.NewMeterProvider(
		metric.WithView(view),
	)

	// Below is an example of how the view will
	// function in the SDK for certain instruments.
	stream, _ := view(metric.Instrument{
		Name: "computation.time.ms",
		Unit: "1",
	})
	fmt.Println("name:", stream.Name)
	fmt.Println("unit:", stream.Unit)
}
