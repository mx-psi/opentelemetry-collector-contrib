// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build js

package opampextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"

import (
	"errors"
	"os"
)

func sendSIGHUP(_ *os.Process) error {
	return errors.New("restart via SIGHUP not supported on this platform")
}
