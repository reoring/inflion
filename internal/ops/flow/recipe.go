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

import "errors"

type Recipe struct {
	Conditions Conditions
	Stages     Stages
}

func (r Recipe) Resolve() (Recipe, error) {
	for ci, c := range r.Conditions.Conditions {
		newCondition, err := c.ResolveNextStage(r)
		if err != nil {
			return r, err
		}
		r.Conditions.Conditions[ci] = newCondition
	}

	for si, s := range r.Stages.Stages {
		newStage, err := s.ResolveNextStage(r)
		if err != nil {
			return r, err
		}
		r.Stages.Stages[si] = newStage
	}
	return r, nil
}

func (r Recipe) SearchNodeById(id string) (Node, error) {
	for _, c := range r.Conditions.Conditions {
		if c.Id == id {
			return c, nil
		}
	}

	for _, s := range r.Stages.Stages {
		if s.Id == id {
			return s, nil
		} else if id == "__end__" {
			return Stage{
				Id:        "end",
				Name:      "end",
				NextStage: NextStage{},
				Actions:   Actions{},
			}, nil
		}
	}

	return nil, errors.New(id + " does not exists")
}

