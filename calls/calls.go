package calls

import (
	"fmt"
	"time"
)

type Request struct {
	UserAgent  string
	Data       string
	Time       time.Time
	HTTPMethod string
}

type Call struct {
	Name       string
	Index      int
	ID         string
	Unanswered bool
	Requests   []Request
	LastPhone  time.Time
}

var Calls []Call

func GetCall(id string) (Call, error) {
	for _, c := range Calls {
		if c.ID == id {
			return c, nil
		}
	}

	return Call{}, fmt.Errorf("unimplemented")
}

func GetCallByIndex(index int) (Call, error) {
	for _, c := range Calls {
		if c.Index == index {
			return c, nil
		}
	}

	return Call{}, fmt.Errorf("unimplemented")
}

func UpdateByID(id string, call Call) {
	for i, c := range Calls {
		if c.ID == id {
			Calls[i] = call
		}
	}
}
