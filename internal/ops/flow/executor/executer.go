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
)

type Executor interface {
	Run(r flow.Recipe) (ExecutionResult, error)
}

type ExecutionResult struct {
	Message string
	Items   []ActionResult
}

type DefaultExecutor struct {
	loader ActionLoader
}

func NewExecutor(loader ActionLoader) Executor {
	return DefaultExecutor{loader: loader}
}

func (e DefaultExecutor) Run(r flow.Recipe) (ExecutionResult, error) {
	var items []ActionResult

	for _, s := range r.Stages.Stages {
		for _, a := range s.Actions.Actions {
			e, err := e.loader.Load(a)
			if err != nil {
				return ExecutionResult{Message: "failed"}, err
			}

			r, err := e.Run(a)
			if err != nil {
				return ExecutionResult{Message: "failed"}, err
			}

			items = append(items, r)
		}
	}

	return ExecutionResult{Message: "success", Items: items}, nil
}
