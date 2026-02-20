// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build js

package datadogreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/metadata"
)

// NewFactory creates a factory for DataDog receiver (stub for js/wasm).
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
		receiver.WithTraces(createTracesReceiver, metadata.TracesStability),
		receiver.WithLogs(createLogsReceiver, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTracesReceiver(_ context.Context, _ receiver.Settings, _ component.Config, _ consumer.Traces) (receiver.Traces, error) {
	return nil, errors.New("datadogreceiver is not supported on js/wasm")
}

func createMetricsReceiver(_ context.Context, _ receiver.Settings, _ component.Config, _ consumer.Metrics) (receiver.Metrics, error) {
	return nil, errors.New("datadogreceiver is not supported on js/wasm")
}

func createLogsReceiver(_ context.Context, _ receiver.Settings, _ component.Config, _ consumer.Logs) (receiver.Logs, error) {
	return nil, errors.New("datadogreceiver is not supported on js/wasm")
}
