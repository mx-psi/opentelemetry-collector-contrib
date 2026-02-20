//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datadogsemanticsprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/datadogsemanticsprocessor"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/datadogsemanticsprocessor/internal/metadata"
)

// NewFactory returns a new factory for the Datadog semantics processor (stub for js/wasm).
//
// Deprecated: [v0.146.0] This processor is deprecated and will be removed in a future release.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		metadata.Type,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, metadata.TracesStability),
	)
}

func createDefaultConfig() component.Config {
	return &struct{}{}
}

func createTracesProcessor(_ context.Context, _ processor.Settings, _ component.Config, _ consumer.Traces) (processor.Traces, error) {
	return nil, errors.New("datadogsemanticsprocessor is not supported on js/wasm")
}
