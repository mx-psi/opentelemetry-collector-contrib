//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package purefbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver/internal/metadata"
)

// NewFactory creates a factory for Pure Storage FlashBlade receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability))
}

// Config defines configuration for the purefb receiver.
type Config struct{}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createMetricsReceiver(
	_ context.Context,
	_ receiver.Settings,
	_ component.Config,
	_ consumer.Metrics,
) (receiver.Metrics, error) {
	return nil, errors.New("purefb receiver is not supported on js/wasm")
}
