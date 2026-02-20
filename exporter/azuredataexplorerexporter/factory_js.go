//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuredataexplorerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter/internal/metadata"
)

// NewFactory creates a factory for the ADX Exporter (stub for js/wasm).
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &struct{}{}
}

func createMetricsExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Metrics, error) {
	return nil, errors.New("azuredataexplorerexporter is not supported on js/wasm")
}

func createTracesExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Traces, error) {
	return nil, errors.New("azuredataexplorerexporter is not supported on js/wasm")
}

func createLogsExporter(_ context.Context, _ exporter.Settings, _ component.Config) (exporter.Logs, error) {
	return nil, errors.New("azuredataexplorerexporter is not supported on js/wasm")
}
