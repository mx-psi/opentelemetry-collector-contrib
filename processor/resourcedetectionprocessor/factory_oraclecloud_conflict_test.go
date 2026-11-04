// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build omit_detector_oraclecloud && enable_detector_oraclecloud

package resourcedetectionprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/processor/processortest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/metadata"
)

func TestUnifiedTagsOmitTakesPrecedence(t *testing.T) {
	assert.Nil(t, detectorFactories()[internal.DetectorType("oracle_cloud")])
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig().(*Config)
	cfg.Detectors = []string{"oracle_cloud"}
	processor, err := factory.CreateTraces(t.Context(), processortest.NewNopSettings(metadata.Type), cfg, consumertest.NewNop())
	require.ErrorContains(t, err, `detector "oracle_cloud" is not compiled into this binary`)
	assert.Nil(t, processor)
}
