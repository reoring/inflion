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

type Stages struct {
	Stages []Stage
}

type Stage struct {
	Id        string
	Name      string
	NextStage NextStage
	Actions   Actions
}

func (s Stages) Get(index int) Stage {
	if index < len(s.Stages) {
		return s.Stages[index]
	}

	return Stage{}
}

func (s Stage) GetId() string {
	return s.Id
}
func (s Stage) IsEmpty() bool {
	return s.Id == ""
}

func (s Stage) ResolveNextStage(r Recipe) (Stage, error) {
	node, err := r.SearchNodeById(s.NextStage.Id)
	if err != nil {
		return Stage{}, err
	}

	return Stage{
		Id:   s.Id,
		Name: s.Name,
		NextStage: NextStage{
			Id:   s.NextStage.Id,
			Node: &node,
		},
		Actions: s.Actions,
	}, nil
}
