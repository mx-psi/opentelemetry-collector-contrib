// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build js

package dbstorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"

import (
	"context"
	"database/sql"
	"errors"

	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"
)

func newClient(_ context.Context, _ *zap.Logger, _ *sql.DB, _, _ string) (*dbStorageClient, error) {
	return nil, errors.New("dbstorage is not supported on js/wasm")
}

type dbStorageClient struct{}

func (c *dbStorageClient) Get(_ context.Context, _ string) ([]byte, error) {
	return nil, errors.New("dbstorage is not supported on js/wasm")
}

func (c *dbStorageClient) Set(_ context.Context, _ string, _ []byte) error {
	return errors.New("dbstorage is not supported on js/wasm")
}

func (c *dbStorageClient) Delete(_ context.Context, _ string) error {
	return errors.New("dbstorage is not supported on js/wasm")
}

func (c *dbStorageClient) Batch(_ context.Context, _ ...*storage.Operation) error {
	return errors.New("dbstorage is not supported on js/wasm")
}

func (c *dbStorageClient) Close(_ context.Context) error {
	return errors.New("dbstorage is not supported on js/wasm")
}
