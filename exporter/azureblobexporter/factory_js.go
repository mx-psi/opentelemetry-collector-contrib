//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azureblobexporter"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azureblobexporter/internal/metadata"
)

// NewFactory creates a factory for Azure Blob exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
		exporter.WithLogs(createLogsExporter, metadata.LogsStability),
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Auth: Authentication{
			Type: ConnectionString,
		},
		Container: TelemetryConfig{
			Metrics: "metrics",
			Logs:    "logs",
			Traces:  "traces",
		},
		BlobNameFormat: BlobNameFormat{
			MetricsFormat:     "2006/01/02/metrics_15_04_05.json",
			LogsFormat:        "2006/01/02/logs_15_04_05.json",
			TracesFormat:      "2006/01/02/traces_15_04_05.json",
			SerialNumEnabled:  true,
			SerialNumRange:    10000,
			Params:            map[string]string{},
			TemplateEnabled:   false,
			TimeParserEnabled: true,
			TimeParserRanges:  nil,
		},
		FormatType: formatTypeJSON,
		AppendBlob: AppendBlob{
			Enabled:   false,
			Separator: "\n",
		},
		Encodings:     Encodings{},
		BackOffConfig: configretry.NewDefaultBackOffConfig(),
	}
}

func createLogsExporter(_ context.Context,
	_ exporter.Settings,
	_ component.Config,
) (exporter.Logs, error) {
	return nil, errors.New("azureblobexporter is not supported on js/wasm")
}

func createMetricsExporter(_ context.Context,
	_ exporter.Settings,
	_ component.Config,
) (exporter.Metrics, error) {
	return nil, errors.New("azureblobexporter is not supported on js/wasm")
}

func createTracesExporter(_ context.Context,
	_ exporter.Settings,
	_ component.Config,
) (exporter.Traces, error) {
	return nil, errors.New("azureblobexporter is not supported on js/wasm")
}

