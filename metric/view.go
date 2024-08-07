package main

import (
	"fmt"

	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	// Create a view that renames the "latency" instrument from the v0.34.0
	// version of the "http" instrumentation library as "request.latency".
	view := metric.NewView(metric.Instrument{
		Name: "latency",
		Scope: instrumentation.Scope{
			Name:    "http",
			Version: "0.34.0",
		},
	}, metric.Stream{Name: "request.latency"})

	// The created view can then be registered with the OpenTelemetry metric
	// SDK using the WithView option.
	_ = metric.NewMeterProvider(
		metric.WithView(view),
	)

	// Below is an example of how the view will
	// function in the SDK for certain instruments.
	stream, _ := view(metric.Instrument{
		Name:        "latency",
		Description: "request latency",
		Unit:        "ms",
		Kind:        metric.InstrumentKindCounter,
		Scope: instrumentation.Scope{
			Name:      "http",
			Version:   "0.34.0",
			SchemaURL: "https://opentelemetry.io/schemas/1.0.0",
		},
	})
	fmt.Println("name:", stream.Name)
	fmt.Println("description:", stream.Description)
	fmt.Println("unit:", stream.Unit)
}
