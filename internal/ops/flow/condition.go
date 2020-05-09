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

type Conditions struct {
	Conditions []Condition
}

func (c Conditions) Get(index int) Condition {
	if index < len(c.Conditions) {
		return c.Conditions[index]
	}

	return Condition{}
}

type Condition struct {
	Id          string
	Expressions []Expression
	IfTrue      NextStage
	IfFalse     NextStage
}

func (c Condition) GetId() string {
	return c.Id
}
func (c Condition) IsEmpty() bool {
	return c.Id == ""
}

func (c Condition) ResolveNextStage(r Recipe) (Condition, error) {
	trueNode, trueErr := r.SearchNodeById(c.IfTrue.Id)
	falseNode, falseErr := r.SearchNodeById(c.IfFalse.Id)

	if trueErr != nil {
		return Condition{}, trueErr
	}
	if falseErr != nil {
		return Condition{}, falseErr
	}

	return Condition{
		Id:          c.Id,
		Expressions: c.Expressions,
		IfTrue: NextStage{
			Id:   c.IfTrue.Id,
			Node: &trueNode,
		},
		IfFalse: NextStage{
			Id:   c.IfFalse.Id,
			Node: &falseNode,
		},
	}, nil
}

type Expression struct {
	Input     string
	Operation string
	Value     string
}
