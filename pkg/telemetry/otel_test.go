// Copyright The Linux Foundation and each contributor to LFX.
// SPDX-License-Identifier: MIT

package telemetry

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSetupOTelSDKWithConfig_NoSchemaURLConflict guards against the semconv
// import drifting from the version used by the OTel SDK's resource.Default(),
// which makes resource.Merge fail with a conflicting Schema URL.
func TestSetupOTelSDKWithConfig_NoSchemaURLConflict(t *testing.T) {
	t.Setenv("OTEL_TRACES_EXPORTER", OTelExporterNone)
	t.Setenv("OTEL_METRICS_EXPORTER", OTelExporterNone)
	t.Setenv("OTEL_LOGS_EXPORTER", OTelExporterNone)

	ctx := context.Background()
	cfg := OTelConfigFromEnv()

	_, err := newResource(cfg)
	require.NoError(t, err)

	shutdown, err := SetupOTelSDKWithConfig(ctx, cfg)
	require.NoError(t, err)
	require.NoError(t, shutdown(ctx))
}
