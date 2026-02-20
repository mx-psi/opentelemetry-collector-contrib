//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package resourcedetectionprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/xprocessor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/metadata"
)

// Config defines configuration for Resource Detection processor.
type Config struct{}

// NewFactory creates a new factory for ResourceDetection processor.
func NewFactory() processor.Factory {
	return xprocessor.NewFactory(
		metadata.Type,
		createDefaultConfig,
		xprocessor.WithTraces(createTracesProcessor, metadata.TracesStability),
		xprocessor.WithMetrics(createMetricsProcessor, metadata.MetricsStability),
		xprocessor.WithLogs(createLogsProcessor, metadata.LogsStability),
		xprocessor.WithProfiles(createProfilesProcessor, metadata.ProfilesStability),
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
	return nil, errors.New("resourcedetection processor is not supported on js/wasm")
}

func createMetricsProcessor(
	_ context.Context,
	_ processor.Settings,
	_ component.Config,
	_ consumer.Metrics,
) (processor.Metrics, error) {
	return nil, errors.New("resourcedetection processor is not supported on js/wasm")
}

func createLogsProcessor(
	_ context.Context,
	_ processor.Settings,
	_ component.Config,
	_ consumer.Logs,
) (processor.Logs, error) {
	return nil, errors.New("resourcedetection processor is not supported on js/wasm")
}

func createProfilesProcessor(
	_ context.Context,
	_ processor.Settings,
	_ component.Config,
	_ xconsumer.Profiles,
) (xprocessor.Profiles, error) {
	return nil, errors.New("resourcedetection processor is not supported on js/wasm")
}
