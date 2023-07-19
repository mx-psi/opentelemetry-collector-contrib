// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbCacheOperationsDataPoint(ts, 1, AttributeTypeHit)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbCollectionCountDataPoint(ts, 1, "database-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbConnectionCountDataPoint(ts, 1, "database-val", AttributeConnectionTypeActive)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbCursorCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbCursorTimeoutCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbDataSizeDataPoint(ts, 1, "database-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbDatabaseCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbDocumentOperationCountDataPoint(ts, 1, "database-val", AttributeOperationInsert)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbExtentCountDataPoint(ts, 1, "database-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbGlobalLockTimeDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordMongodbHealthDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbIndexAccessCountDataPoint(ts, 1, "database-val", "collection-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbIndexCountDataPoint(ts, 1, "database-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbIndexSizeDataPoint(ts, 1, "database-val")

			allMetricsCount++
			mb.RecordMongodbLockAcquireCountDataPoint(ts, 1, "database-val", AttributeLockTypeParallelBatchWriteMode, AttributeLockModeShared)

			allMetricsCount++
			mb.RecordMongodbLockAcquireTimeDataPoint(ts, 1, "database-val", AttributeLockTypeParallelBatchWriteMode, AttributeLockModeShared)

			allMetricsCount++
			mb.RecordMongodbLockAcquireWaitCountDataPoint(ts, 1, "database-val", AttributeLockTypeParallelBatchWriteMode, AttributeLockModeShared)

			allMetricsCount++
			mb.RecordMongodbLockDeadlockCountDataPoint(ts, 1, "database-val", AttributeLockTypeParallelBatchWriteMode, AttributeLockModeShared)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbMemoryUsageDataPoint(ts, 1, "database-val", AttributeMemoryTypeResident)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbNetworkIoReceiveDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbNetworkIoTransmitDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbNetworkRequestCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbObjectCountDataPoint(ts, 1, "database-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbOperationCountDataPoint(ts, 1, AttributeOperationInsert)

			allMetricsCount++
			mb.RecordMongodbOperationLatencyTimeDataPoint(ts, 1, AttributeOperationLatencyRead)

			allMetricsCount++
			mb.RecordMongodbOperationReplCountDataPoint(ts, 1, AttributeOperationInsert)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbOperationTimeDataPoint(ts, 1, AttributeOperationInsert)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbSessionCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordMongodbStorageSizeDataPoint(ts, 1, "database-val")

			allMetricsCount++
			mb.RecordMongodbUptimeDataPoint(ts, 1)

			metrics := mb.Emit(WithDatabase("database-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("database")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.Database.Enabled, ok)
			if mb.resourceAttributesConfig.Database.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "database-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 1)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "mongodb.cache.operations":
					assert.False(t, validatedMetrics["mongodb.cache.operations"], "Found a duplicate in the metrics slice: mongodb.cache.operations")
					validatedMetrics["mongodb.cache.operations"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of cache operations of the instance.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "hit", attrVal.Str())
				case "mongodb.collection.count":
					assert.False(t, validatedMetrics["mongodb.collection.count"], "Found a duplicate in the metrics slice: mongodb.collection.count")
					validatedMetrics["mongodb.collection.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of collections.", ms.At(i).Description())
					assert.Equal(t, "{collections}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
				case "mongodb.connection.count":
					assert.False(t, validatedMetrics["mongodb.connection.count"], "Found a duplicate in the metrics slice: mongodb.connection.count")
					validatedMetrics["mongodb.connection.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of connections.", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "active", attrVal.Str())
				case "mongodb.cursor.count":
					assert.False(t, validatedMetrics["mongodb.cursor.count"], "Found a duplicate in the metrics slice: mongodb.cursor.count")
					validatedMetrics["mongodb.cursor.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of open cursors maintained for clients.", ms.At(i).Description())
					assert.Equal(t, "{cursors}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.cursor.timeout.count":
					assert.False(t, validatedMetrics["mongodb.cursor.timeout.count"], "Found a duplicate in the metrics slice: mongodb.cursor.timeout.count")
					validatedMetrics["mongodb.cursor.timeout.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of cursors that have timed out.", ms.At(i).Description())
					assert.Equal(t, "{cursors}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.data.size":
					assert.False(t, validatedMetrics["mongodb.data.size"], "Found a duplicate in the metrics slice: mongodb.data.size")
					validatedMetrics["mongodb.data.size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The size of the collection. Data compression does not affect this value.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
				case "mongodb.database.count":
					assert.False(t, validatedMetrics["mongodb.database.count"], "Found a duplicate in the metrics slice: mongodb.database.count")
					validatedMetrics["mongodb.database.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of existing databases.", ms.At(i).Description())
					assert.Equal(t, "{databases}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.document.operation.count":
					assert.False(t, validatedMetrics["mongodb.document.operation.count"], "Found a duplicate in the metrics slice: mongodb.document.operation.count")
					validatedMetrics["mongodb.document.operation.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of document operations executed.", ms.At(i).Description())
					assert.Equal(t, "{documents}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.EqualValues(t, "insert", attrVal.Str())
				case "mongodb.extent.count":
					assert.False(t, validatedMetrics["mongodb.extent.count"], "Found a duplicate in the metrics slice: mongodb.extent.count")
					validatedMetrics["mongodb.extent.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of extents.", ms.At(i).Description())
					assert.Equal(t, "{extents}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
				case "mongodb.global_lock.time":
					assert.False(t, validatedMetrics["mongodb.global_lock.time"], "Found a duplicate in the metrics slice: mongodb.global_lock.time")
					validatedMetrics["mongodb.global_lock.time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The time the global lock has been held.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.health":
					assert.False(t, validatedMetrics["mongodb.health"], "Found a duplicate in the metrics slice: mongodb.health")
					validatedMetrics["mongodb.health"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The health status of the server.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.index.access.count":
					assert.False(t, validatedMetrics["mongodb.index.access.count"], "Found a duplicate in the metrics slice: mongodb.index.access.count")
					validatedMetrics["mongodb.index.access.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of times an index has been accessed.", ms.At(i).Description())
					assert.Equal(t, "{accesses}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("collection")
					assert.True(t, ok)
					assert.EqualValues(t, "collection-val", attrVal.Str())
				case "mongodb.index.count":
					assert.False(t, validatedMetrics["mongodb.index.count"], "Found a duplicate in the metrics slice: mongodb.index.count")
					validatedMetrics["mongodb.index.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of indexes.", ms.At(i).Description())
					assert.Equal(t, "{indexes}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
				case "mongodb.index.size":
					assert.False(t, validatedMetrics["mongodb.index.size"], "Found a duplicate in the metrics slice: mongodb.index.size")
					validatedMetrics["mongodb.index.size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Sum of the space allocated to all indexes in the database, including free index space.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
				case "mongodb.lock.acquire.count":
					assert.False(t, validatedMetrics["mongodb.lock.acquire.count"], "Found a duplicate in the metrics slice: mongodb.lock.acquire.count")
					validatedMetrics["mongodb.lock.acquire.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times the lock was acquired in the specified mode.", ms.At(i).Description())
					assert.Equal(t, "{count}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_type")
					assert.True(t, ok)
					assert.EqualValues(t, "parallel_batch_write_mode", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_mode")
					assert.True(t, ok)
					assert.EqualValues(t, "shared", attrVal.Str())
				case "mongodb.lock.acquire.time":
					assert.False(t, validatedMetrics["mongodb.lock.acquire.time"], "Found a duplicate in the metrics slice: mongodb.lock.acquire.time")
					validatedMetrics["mongodb.lock.acquire.time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Cumulative wait time for the lock acquisitions.", ms.At(i).Description())
					assert.Equal(t, "microseconds", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_type")
					assert.True(t, ok)
					assert.EqualValues(t, "parallel_batch_write_mode", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_mode")
					assert.True(t, ok)
					assert.EqualValues(t, "shared", attrVal.Str())
				case "mongodb.lock.acquire.wait_count":
					assert.False(t, validatedMetrics["mongodb.lock.acquire.wait_count"], "Found a duplicate in the metrics slice: mongodb.lock.acquire.wait_count")
					validatedMetrics["mongodb.lock.acquire.wait_count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times the lock acquisitions encountered waits because the locks were held in a conflicting mode.", ms.At(i).Description())
					assert.Equal(t, "{count}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_type")
					assert.True(t, ok)
					assert.EqualValues(t, "parallel_batch_write_mode", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_mode")
					assert.True(t, ok)
					assert.EqualValues(t, "shared", attrVal.Str())
				case "mongodb.lock.deadlock.count":
					assert.False(t, validatedMetrics["mongodb.lock.deadlock.count"], "Found a duplicate in the metrics slice: mongodb.lock.deadlock.count")
					validatedMetrics["mongodb.lock.deadlock.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times the lock acquisitions encountered deadlocks.", ms.At(i).Description())
					assert.Equal(t, "{count}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_type")
					assert.True(t, ok)
					assert.EqualValues(t, "parallel_batch_write_mode", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("lock_mode")
					assert.True(t, ok)
					assert.EqualValues(t, "shared", attrVal.Str())
				case "mongodb.memory.usage":
					assert.False(t, validatedMetrics["mongodb.memory.usage"], "Found a duplicate in the metrics slice: mongodb.memory.usage")
					validatedMetrics["mongodb.memory.usage"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of memory used.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.EqualValues(t, "resident", attrVal.Str())
				case "mongodb.network.io.receive":
					assert.False(t, validatedMetrics["mongodb.network.io.receive"], "Found a duplicate in the metrics slice: mongodb.network.io.receive")
					validatedMetrics["mongodb.network.io.receive"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of bytes received.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.network.io.transmit":
					assert.False(t, validatedMetrics["mongodb.network.io.transmit"], "Found a duplicate in the metrics slice: mongodb.network.io.transmit")
					validatedMetrics["mongodb.network.io.transmit"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of by transmitted.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.network.request.count":
					assert.False(t, validatedMetrics["mongodb.network.request.count"], "Found a duplicate in the metrics slice: mongodb.network.request.count")
					validatedMetrics["mongodb.network.request.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of requests received by the server.", ms.At(i).Description())
					assert.Equal(t, "{requests}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.object.count":
					assert.False(t, validatedMetrics["mongodb.object.count"], "Found a duplicate in the metrics slice: mongodb.object.count")
					validatedMetrics["mongodb.object.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of objects.", ms.At(i).Description())
					assert.Equal(t, "{objects}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
				case "mongodb.operation.count":
					assert.False(t, validatedMetrics["mongodb.operation.count"], "Found a duplicate in the metrics slice: mongodb.operation.count")
					validatedMetrics["mongodb.operation.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of operations executed.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.EqualValues(t, "insert", attrVal.Str())
				case "mongodb.operation.latency.time":
					assert.False(t, validatedMetrics["mongodb.operation.latency.time"], "Found a duplicate in the metrics slice: mongodb.operation.latency.time")
					validatedMetrics["mongodb.operation.latency.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The latency of operations.", ms.At(i).Description())
					assert.Equal(t, "us", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.EqualValues(t, "read", attrVal.Str())
				case "mongodb.operation.repl.count":
					assert.False(t, validatedMetrics["mongodb.operation.repl.count"], "Found a duplicate in the metrics slice: mongodb.operation.repl.count")
					validatedMetrics["mongodb.operation.repl.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of replicated operations executed.", ms.At(i).Description())
					assert.Equal(t, "{operations}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.EqualValues(t, "insert", attrVal.Str())
				case "mongodb.operation.time":
					assert.False(t, validatedMetrics["mongodb.operation.time"], "Found a duplicate in the metrics slice: mongodb.operation.time")
					validatedMetrics["mongodb.operation.time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total time spent performing operations.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.EqualValues(t, "insert", attrVal.Str())
				case "mongodb.session.count":
					assert.False(t, validatedMetrics["mongodb.session.count"], "Found a duplicate in the metrics slice: mongodb.session.count")
					validatedMetrics["mongodb.session.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of active sessions.", ms.At(i).Description())
					assert.Equal(t, "{sessions}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "mongodb.storage.size":
					assert.False(t, validatedMetrics["mongodb.storage.size"], "Found a duplicate in the metrics slice: mongodb.storage.size")
					validatedMetrics["mongodb.storage.size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total amount of storage allocated to this collection.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "database-val", attrVal.Str())
				case "mongodb.uptime":
					assert.False(t, validatedMetrics["mongodb.uptime"], "Found a duplicate in the metrics slice: mongodb.uptime")
					validatedMetrics["mongodb.uptime"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of time that the server has been running.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
