// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package executor

import (
	"github.com/inflion/inflion/internal/ops/flow"
	"testing"
)

var recipe flow.Recipe

func init() {
	recipe = flow.Reader{}.Read()
}

func TestExecutor(t *testing.T) {
	r, err := recipe.Resolve()
	if err != nil {
		t.Error(err)
	}

	re := NewExecutor(AggregateActionLoader{})
	result, err := re.Run(r)
	if err != nil {
		t.Error(err)
	}

	got := result.Items[0].Outputs["result"]
	want := "ok"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
