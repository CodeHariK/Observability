package main

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	// Create a view that removes the "http.request.method" attribute recorded
	// by the "latency" instrument from the "http" instrumentation library.
	view := metric.NewView(
		metric.Instrument{
			Name:  "latency",
			Scope: instrumentation.Scope{Name: "http"},
		},
		metric.Stream{AttributeFilter: attribute.NewDenyKeysFilter("http.request.method")},
	)

	// The created view can then be registered with the OpenTelemetry metric
	// SDK using the WithView option.
	_ = metric.NewMeterProvider(
		metric.WithView(view),
	)
}
