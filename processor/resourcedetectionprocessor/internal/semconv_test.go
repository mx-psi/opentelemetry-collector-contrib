// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/testutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/metadata"
)

func TestSemconvOSType(t *testing.T) {
	for _, tc := range []struct {
		name     string
		enabled  bool
		osType   string
		expected string
	}{
		{name: "gate disabled", enabled: false, osType: "z_os", expected: "z_os"},
		{name: "gate enabled", enabled: true, osType: "z_os", expected: "zos"},
		{name: "gate enabled, other value", enabled: true, osType: "linux", expected: "linux"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer testutil.SetFeatureGateForTest(t, metadata.ProcessorResourceDetectionSemconvValueFixesFeatureGate, tc.enabled)()
			assert.Equal(t, tc.expected, SemconvOSType(tc.osType))
		})
	}
}

func TestSemconvIntAttribute(t *testing.T) {
	for _, tc := range []struct {
		name     string
		enabled  bool
		attrs    map[string]any
		expected map[string]any
	}{
		{name: "gate disabled", enabled: false, attrs: map[string]any{"k": "128"}, expected: map[string]any{"k": "128"}},
		{name: "gate enabled", enabled: true, attrs: map[string]any{"k": "128"}, expected: map[string]any{"k": int64(128)}},
		{name: "gate enabled, not an int", enabled: true, attrs: map[string]any{"k": "abc"}, expected: map[string]any{"k": "abc"}},
		{name: "gate enabled, missing", enabled: true, attrs: map[string]any{}, expected: map[string]any{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer testutil.SetFeatureGateForTest(t, metadata.ProcessorResourceDetectionSemconvValueFixesFeatureGate, tc.enabled)()
			attrs := pcommon.NewMap()
			assert.NoError(t, attrs.FromRaw(tc.attrs))
			SemconvIntAttribute(attrs, "k")
			assert.Equal(t, tc.expected, attrs.AsRaw())
		})
	}
}
