// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !js

package opampextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"

import (
	"os"
	"syscall"
)

func sendSIGHUP(p *os.Process) error {
	return p.Signal(syscall.SIGHUP)
}
