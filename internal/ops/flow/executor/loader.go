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
	"errors"
	"github.com/inflion/inflion/internal/ops/flow"
)

type ActionLoader interface {
	Load(action flow.Action) (ActionExecutor, error)
}

type AggregateActionLoader struct {
}

func (a AggregateActionLoader) Load(action flow.Action) (ActionExecutor, error) {
	loaders := []ActionLoader{
		EmbeddedActionLoader{},
	}

	for _, l := range loaders {
		a, err := l.Load(action)
		if err != nil {
			return a, err
		}

		return a, nil
	}

	return NullActionExecutor{}, errors.New("no action found at type: " + action.Type)
}

type EmbeddedActionLoader struct {
}

func (e EmbeddedActionLoader) Load(action flow.Action) (ActionExecutor, error) {
	switch action.Type {
	case "matcher":
		return MatcherActionExecutor{}, nil
	}

	return NullActionExecutor{}, errors.New("no embedded action found at " + action.Type)
}

