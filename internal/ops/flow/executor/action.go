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

import "github.com/inflion/inflion/internal/ops/flow"

type ActionResult struct {
	Action  flow.Action
	Outputs map[string]string
}

type ActionExecutor interface {
	Run(action flow.Action) (ActionResult, error)
}

type NullActionExecutor struct {
}

func (n NullActionExecutor) Run(action flow.Action) (ActionResult, error) {
	return ActionResult{}, nil
}

type MatcherActionExecutor struct {
}

func (m MatcherActionExecutor) Run(action flow.Action) (ActionResult, error) {
	return ActionResult{
		Action: action,
		Outputs: map[string]string{
			"result": "ok",
		},
	}, nil
}

