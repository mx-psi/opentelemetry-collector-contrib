//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filestorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage/internal/metadata"
)

// NewFactory creates a factory for HostObserver extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Directory: "/var/lib/otelcol/file_storage",
		Compaction: &CompactionConfig{
			Directory:                  "/var/lib/otelcol/file_storage",
			OnStart:                    false,
			OnRebound:                  false,
			MaxTransactionSize:         65536,
			ReboundNeededThresholdMiB:  100,
			ReboundTriggerThresholdMiB: 10,
			CheckInterval:              time.Second * 5,
			CleanupOnStart:             false,
		},
		Timeout:              time.Second,
		FSync:                false,
		CreateDirectory:      false,
		DirectoryPermissions: "0750",
	}
}

func createExtension(
	_ context.Context,
	_ extension.Settings,
	_ component.Config,
) (extension.Extension, error) {
	return nil, errors.New("filestorage extension is not supported on js/wasm")
}
