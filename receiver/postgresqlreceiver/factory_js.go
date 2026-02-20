//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package postgresqlreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver/internal/metadata"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
		receiver.WithLogs(createLogsReceiver, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	cfg := scraperhelper.NewDefaultControllerConfig()
	cfg.CollectionInterval = 10 * time.Second

	return &Config{
		ControllerConfig: cfg,
		AddrConfig: confignet.AddrConfig{
			Endpoint:  "localhost:5432",
			Transport: confignet.TransportTypeTCP,
		},
		ClientConfig: configtls.ClientConfig{
			Insecure:           false,
			InsecureSkipVerify: true,
		},
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
		LogsBuilderConfig:    metadata.DefaultLogsBuilderConfig(),
		QuerySampleCollection: QuerySampleCollection{
			MaxRowsPerQuery: 1000,
		},
		TopQueryCollection: TopQueryCollection{
			TopNQuery:              200,
			MaxRowsPerQuery:        1000,
			MaxExplainEachInterval: 1000,
			QueryPlanCacheSize:     1000,
			QueryPlanCacheTTL:      time.Hour,
		},
	}
}

func createMetricsReceiver(
	_ context.Context,
	_ receiver.Settings,
	_ component.Config,
	_ consumer.Metrics,
) (receiver.Metrics, error) {
	return nil, errors.New("postgresqlreceiver is not supported on js/wasm")
}

func createLogsReceiver(
	_ context.Context,
	_ receiver.Settings,
	_ component.Config,
	_ consumer.Logs,
) (receiver.Logs, error) {
	return nil, errors.New("postgresqlreceiver is not supported on js/wasm")
}
