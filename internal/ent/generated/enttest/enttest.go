// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package enttest

import (
	"context"

	"go.infratographer.com/virtual-machine-api/internal/ent/generated"
	// required by schema hooks.
	_ "go.infratographer.com/virtual-machine-api/internal/ent/generated/runtime"

	"entgo.io/ent/dialect/sql/schema"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/migrate"
)

type (
	// TestingT is the interface that is shared between
	// testing.T and testing.B and used by enttest.
	TestingT interface {
		FailNow()
		Error(...any)
	}

	// Option configures client creation.
	Option func(*options)

	options struct {
		opts        []generated.Option
		migrateOpts []schema.MigrateOption
	}
)

// WithOptions forwards options to client creation.
func WithOptions(opts ...generated.Option) Option {
	return func(o *options) {
		o.opts = append(o.opts, opts...)
	}
}

// WithMigrateOptions forwards options to auto migration.
func WithMigrateOptions(opts ...schema.MigrateOption) Option {
	return func(o *options) {
		o.migrateOpts = append(o.migrateOpts, opts...)
	}
}

func newOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Open calls generated.Open and auto-run migration.
func Open(t TestingT, driverName, dataSourceName string, opts ...Option) *generated.Client {
	o := newOptions(opts)
	c, err := generated.Open(driverName, dataSourceName, o.opts...)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	migrateSchema(t, c, o)
	return c
}

// NewClient calls generated.NewClient and auto-run migration.
func NewClient(t TestingT, opts ...Option) *generated.Client {
	o := newOptions(opts)
	c := generated.NewClient(o.opts...)
	migrateSchema(t, c, o)
	return c
}
func migrateSchema(t TestingT, c *generated.Client, o *options) {
	tables, err := schema.CopyTables(migrate.Tables)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := migrate.Create(context.Background(), c.Schema, tables, o.migrateOpts...); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
