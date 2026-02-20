//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redactionprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor/internal/metadata"
)

// NewFactory creates a factory for the redaction processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		metadata.Type,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, metadata.TracesStability),
		processor.WithLogs(createLogsProcessor, metadata.LogsStability),
		processor.WithMetrics(createMetricsProcessor, metadata.MetricsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTracesProcessor(
	_ context.Context,
	_ processor.Settings,
	_ component.Config,
	_ consumer.Traces,
) (processor.Traces, error) {
	return nil, errors.New("redactionprocessor is not supported on js/wasm")
}

func createLogsProcessor(
	_ context.Context,
	_ processor.Settings,
	_ component.Config,
	_ consumer.Logs,
) (processor.Logs, error) {
	return nil, errors.New("redactionprocessor is not supported on js/wasm")
}

func createMetricsProcessor(
	_ context.Context,
	_ processor.Settings,
	_ component.Config,
	_ consumer.Metrics,
) (processor.Metrics, error) {
	return nil, errors.New("redactionprocessor is not supported on js/wasm")
}
