package main

import (
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	// Create a view that makes the "latency" instrument from the "http"
	// instrumentation library to be reported as an exponential histogram.
	view := metric.NewView(
		metric.Instrument{
			Name:  "latency",
			Scope: instrumentation.Scope{Name: "http"},
		},
		metric.Stream{
			Aggregation: metric.AggregationBase2ExponentialHistogram{
				MaxSize:  160,
				MaxScale: 20,
			},
		},
	)

	// The created view can then be registered with the OpenTelemetry metric
	// SDK using the WithView option.
	_ = metric.NewMeterProvider(
		metric.WithView(view),
	)
}
