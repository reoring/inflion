// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package flow

import (
	"testing"
)

var recipe Recipe

func init() {
	recipe = Reader{}.Read()
}

func TestConditionOne(t *testing.T) {
	r, err := recipe.Resolve()
	if err != nil {
		t.Error(err)
	}

	want := "s2"
	got := r.Conditions.Get(0).IfTrue.Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

func TestConditionTwo(t *testing.T) {
	r, err := recipe.Resolve()
	if err != nil {
		t.Error(err)
	}

	want := "s3"
	got := r.Conditions.Get(0).IfFalse.Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

func TestShouldEmptyNoSuchCondition(t *testing.T) {
	r, err := recipe.Resolve()
	if err != nil {
		t.Error(err)
	}

	got := r.Conditions.Get(999).IsEmpty()
	if true != got {
		t.Errorf("got %v want true", got)
	}
}

func Test_BoundaryValue_ShouldNotEmptyCondition(t *testing.T) {
	r, err := recipe.Resolve()
	if err != nil {
		t.Error(err)
	}

	got := r.Conditions.Get(0).IsEmpty()
	if false != got {
		t.Errorf("got %v want false", got)
	}
}

func TestShouldStageS1sNextWillBeCondition1(t *testing.T) {
	r, err := recipe.Resolve()
	if err != nil {
		t.Error(err)
	}

	want := "c1"
	got := r.Stages.Get(0).NextStage.Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

func TestShouldEmptyNoSuchStage(t *testing.T) {
	r, err := recipe.Resolve()
	if err != nil {
		t.Error(err)
	}

	got := r.Stages.Get(999).IsEmpty()
	if true != got {
		t.Errorf("got %v want true", got)
	}
}

func Test_BoundaryValue_ShouldNotEmptyStage(t *testing.T) {
	r, err := recipe.Resolve()
	if err != nil {
		t.Error(err)
	}

	got := r.Stages.Get(2).IsEmpty()
	if false != got {
		t.Errorf("got %v want false", got)
	}
}
