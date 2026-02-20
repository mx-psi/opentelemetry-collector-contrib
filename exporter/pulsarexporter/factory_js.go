//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pulsarexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter/internal/metadata"
)

// Config is a stub for js/wasm builds.
type Config struct{}

func (c *Config) Validate() error { return nil }

// FactoryOption applies changes to pulsarExporterFactory.
type FactoryOption func()

// NewFactory creates a stub Pulsar exporter factory for js/wasm.
func NewFactory(_ ...FactoryOption) exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTracesExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Traces, error) {
	return nil, errors.New("pulsarexporter is not supported on js/wasm")
}

func createMetricsExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Metrics, error) {
	return nil, errors.New("pulsarexporter is not supported on js/wasm")
}

func createLogsExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Logs, error) {
	return nil, errors.New("pulsarexporter is not supported on js/wasm")
}
