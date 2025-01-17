// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

var _ ottl.Path[any] = &TestPath[any]{}

type TestPath[K any] struct {
	N        string
	Keys     ottl.Key[K]
	NextPath *TestPath[K]
}

func (p *TestPath[K]) Name() string {
	return p.N
}

func (p *TestPath[K]) Next() ottl.Path[K] {
	if p.NextPath == nil {
		return nil
	}
	return p.NextPath
}

func (p *TestPath[K]) Key() ottl.Key[K] {
	return p.Keys
}

var _ ottl.Key[any] = &TestKey[any]{}

type TestKey[K any] struct {
	S       *string
	I       *int64
	NextKey *TestKey[K]
}

func (k *TestKey[K]) String(_ context.Context, _ K) (*string, error) {
	return k.S, nil
}

func (k *TestKey[K]) Int(_ context.Context, _ K) (*int64, error) {
	return k.I, nil
}

func (k *TestKey[K]) Next() ottl.Key[K] {
	if k.NextKey == nil {
		return nil
	}
	return k.NextKey
}
