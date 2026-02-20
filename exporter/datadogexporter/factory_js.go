//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datadogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metadata"
)

// NewFactory creates a Datadog exporter factory (stub for js/wasm).
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &struct{}{}
}

func createMetricsExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Metrics, error) {
	return nil, errors.New("datadogexporter is not supported on js/wasm")
}

func createTracesExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Traces, error) {
	return nil, errors.New("datadogexporter is not supported on js/wasm")
}

func createLogsExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Logs, error) {
	return nil, errors.New("datadogexporter is not supported on js/wasm")
}
