package flow

type Condition struct {
	Id         string `json:"id"`
	Conditions []struct {
		Input     string `json:"input"`
		Operation string `json:"operation"`
		Value     string `json:"value"`
	}
	NextActions struct {
		True struct {
			Operation string `json:"operation"`
			NextId    string `json:"next_id"`
		} `json:"true"`
		False struct {
			Operation string `json:"operation"`
			NextId    string `json:"next_id"`
		} `json:"false"`
	} `json:"next_actions"`
}

type Stage struct {
	Id string `json:"id"`
	Next string `json:"next"`
	Name string `json:"name"`
	Actions []struct {
		Type string `json:"type"`
		Params map[string]string `json:"params"`
	}
}

type Flow struct {
	Metadata struct {
		Format struct {
			Version int `json:"version"`
		} `json:"Format"`
	} `json:"Metadata"`

	Body struct {
		Conditions []Condition `json:"conditions"`
		Stages     []Stage     `json:"stages"`
	}
}
