//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datadogextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/metadata"
)

// NewFactory creates a factory for the Datadog extension (stub for js/wasm).
func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}

func createDefaultConfig() component.Config {
	return &struct{}{}
}

func createExtension(_ context.Context, _ extension.Settings, _ component.Config) (extension.Extension, error) {
	return nil, errors.New("datadogextension is not supported on js/wasm")
}
