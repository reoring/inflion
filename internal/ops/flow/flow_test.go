package flow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFlowStruct(t *testing.T) {
	bytes, err := ioutil.ReadFile("sample_flow.json")
	if err != nil {
		fmt.Println(err)
	}

	flow := Flow{}
	err = json.Unmarshal(bytes, &flow)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(flow)
}