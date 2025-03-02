// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/odigos/processor/odigostrafficmetrics/internal/metadata"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"go.opentelemetry.io/collector/component/componenttest"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := componenttest.NewTelemetry()
	tb, err := metadata.NewTelemetryBuilder(testTel.NewTelemetrySettings())
	require.NoError(t, err)
	defer tb.Shutdown()
	tb.OdigosAcceptedLogRecords.Add(context.Background(), 1)
	tb.OdigosAcceptedMetricPoints.Add(context.Background(), 1)
	tb.OdigosAcceptedSpans.Add(context.Background(), 1)
	tb.OdigosLogDataSize.Add(context.Background(), 1)
	tb.OdigosMetricDataSize.Add(context.Background(), 1)
	tb.OdigosTraceDataSize.Add(context.Background(), 1)
	AssertEqualOdigosAcceptedLogRecords(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualOdigosAcceptedMetricPoints(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualOdigosAcceptedSpans(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualOdigosLogDataSize(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualOdigosMetricDataSize(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualOdigosTraceDataSize(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())

	require.NoError(t, testTel.Shutdown(context.Background()))
}
