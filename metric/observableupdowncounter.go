package main

import (
	"context"
	"database/sql"

	"go.opentelemetry.io/otel/metric"
)

func main() {
	// The function registers asynchronous metrics for the provided db.
	// Make sure to unregister metric.Registration before closing the provided db.
	_ = func(db *sql.DB, meter metric.Meter, poolName string) (metric.Registration, error) {
		max, err := meter.Int64ObservableUpDownCounter(
			"db.client.connections.max",
			metric.WithDescription("The maximum number of open connections allowed."),
			metric.WithUnit("{connection}"),
		)
		if err != nil {
			return nil, err
		}

		waitTime, err := meter.Int64ObservableUpDownCounter(
			"db.client.connections.wait_time",
			metric.WithDescription("The time it took to obtain an open connection from the pool."),
			metric.WithUnit("ms"),
		)
		if err != nil {
			return nil, err
		}

		reg, err := meter.RegisterCallback(
			func(_ context.Context, o metric.Observer) error {
				stats := db.Stats()
				o.ObserveInt64(max, int64(stats.MaxOpenConnections))
				o.ObserveInt64(waitTime, int64(stats.WaitDuration))
				return nil
			},
			max,
			waitTime,
		)
		if err != nil {
			return nil, err
		}
		return reg, nil
	}
}
