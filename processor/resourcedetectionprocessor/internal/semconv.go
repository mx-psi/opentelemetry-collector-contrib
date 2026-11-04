// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"

import (
	"strconv"

	"go.opentelemetry.io/collector/pdata/pcommon"
	conventions "go.opentelemetry.io/otel/semconv/v1.43.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/metadata"
)

const deprecatedOSTypeZOS = "z_os"

// SemconvOSType maps deprecated os.type values to their replacements
// when the [metadata.ProcessorResourceDetectionSemconvValueFixesFeatureGate] feature gate is enabled.
func SemconvOSType(osType string) string {
	if metadata.ProcessorResourceDetectionSemconvValueFixesFeatureGate.IsEnabled() && osType == deprecatedOSTypeZOS {
		return conventions.OSTypeZOS.Value.AsString()
	}
	return osType
}

// SemconvIntAttribute converts the string attribute at key to an int
// when the [metadata.ProcessorResourceDetectionSemconvValueFixesFeatureGate] feature gate is enabled. Non-integer values are left unchanged.
func SemconvIntAttribute(attrs pcommon.Map, key string) {
	if !metadata.ProcessorResourceDetectionSemconvValueFixesFeatureGate.IsEnabled() {
		return
	}
	v, ok := attrs.Get(key)
	if !ok || v.Type() != pcommon.ValueTypeStr {
		return
	}
	if i, err := strconv.ParseInt(v.Str(), 10, 64); err == nil {
		v.SetInt(i)
	}
}
