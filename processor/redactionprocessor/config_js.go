//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redactionprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor"

// Config is a stub configuration for js/wasm builds.
// The redactionprocessor is not supported on js/wasm due to dependencies
// that do not support that platform.
type Config struct{}
