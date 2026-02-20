// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build js

package dbstorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func createExtension(
	_ context.Context,
	_ extension.Settings,
	_ component.Config,
) (extension.Extension, error) {
	return nil, errors.New("dbstorage is not supported on js/wasm")
}
