//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pulsarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver/internal/metadata"
)

// Config is a stub for js/wasm builds.
type Config struct{}

func (c *Config) Validate() error { return nil }

// FactoryOption applies changes to PulsarExporterFactory.
type FactoryOption func()

// NewFactory creates a stub Pulsar receiver factory for js/wasm.
func NewFactory(_ ...FactoryOption) receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithTraces(createTracesReceiver, metadata.TracesStability),
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
		receiver.WithLogs(createLogsReceiver, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTracesReceiver(_ context.Context, _ receiver.Settings, _ component.Config, _ consumer.Traces) (receiver.Traces, error) {
	return nil, errors.New("pulsarreceiver is not supported on js/wasm")
}

func createMetricsReceiver(_ context.Context, _ receiver.Settings, _ component.Config, _ consumer.Metrics) (receiver.Metrics, error) {
	return nil, errors.New("pulsarreceiver is not supported on js/wasm")
}

func createLogsReceiver(_ context.Context, _ receiver.Settings, _ component.Config, _ consumer.Logs) (receiver.Logs, error) {
	return nil, errors.New("pulsarreceiver is not supported on js/wasm")
}
