//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datadogconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector/internal/metadata"
)

// NewFactory creates a factory for datadog connector (stub for js/wasm).
func NewFactory() connector.Factory {
	return connector.NewFactory(
		metadata.Type,
		createDefaultConfig,
		connector.WithTracesToMetrics(createTracesToMetrics, metadata.TracesToMetricsStability),
		connector.WithTracesToTraces(createTracesToTraces, metadata.TracesToTracesStability),
	)
}

func createDefaultConfig() component.Config {
	return &struct{}{}
}

func createTracesToMetrics(_ context.Context, _ connector.Settings, _ component.Config, _ consumer.Metrics) (connector.Traces, error) {
	return nil, errors.New("datadogconnector is not supported on js/wasm")
}

func createTracesToTraces(_ context.Context, _ connector.Settings, _ component.Config, _ consumer.Traces) (connector.Traces, error) {
	return nil, errors.New("datadogconnector is not supported on js/wasm")
}
