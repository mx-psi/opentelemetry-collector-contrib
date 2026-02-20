//go:build js

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver/internal"

import (
	"database/sql"
	"errors"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"
)

func NewPool(_ sqlquery.SQLOpenerFunc, _, _ string, _ int) interface {
	DB() (*sql.DB, error)
} {
	return &sqlPool{}
}

type sqlPool struct{}

func (sp *sqlPool) DB() (*sql.DB, error) {
	return nil, errors.New("sqlquery receiver is not supported on js/wasm")
}
