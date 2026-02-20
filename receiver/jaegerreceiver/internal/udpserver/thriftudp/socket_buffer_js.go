// Copyright The OpenTelemetry Authors
// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build js

package thriftudp // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/udpserver/thriftudp"

import (
	"net"
)

// Not supported on js/wasm, so js version just returns nil
func setSocketBuffer(_ *net.UDPConn, _ int) error {
	return nil
}
