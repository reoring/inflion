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

type Reader struct{}

func (Reader) Read() Recipe {
	stageS1 := Stage{
		Id:   "s1",
		Name: "s1",
		NextStage: NextStage{
			Id:   "c1",
			Node: nil,
		},
		Actions: Actions{
			Actions: []Action{
				{
					Type:   "matcher",
					Params: nil,
				},
			},
		},
	}
	stageS2 := Stage{
		Id:   "s2",
		Name: "s2",
		NextStage: NextStage{
			Id:   "s3",
			Node: nil,
		},
		Actions: Actions{},
	}
	stageS3 := Stage{
		Id:   "s3",
		Name: "s3",
		NextStage: NextStage{
			Id:   "__end__",
			Node: nil,
		},
		Actions: Actions{},
	}

	return Recipe{
		Conditions: Conditions{
			Conditions: []Condition{
				{
					Id:          "c1",
					Expressions: nil,
					IfTrue: NextStage{
						Id:   "s2",
						Node: nil,
					},
					IfFalse: NextStage{
						Id:   "s3",
						Node: nil,
					},
				},
			},
		},
		Stages: Stages{
			Stages: []Stage{
				stageS1, stageS2, stageS3,
			},
		},
	}
}

