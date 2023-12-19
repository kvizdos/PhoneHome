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

func (c Call) AverageRequestDelay() time.Duration {
	if len(c.Requests) < 2 {
		return 0
	}

	var totalDelay time.Duration
	var count int

	for i := 1; i < len(c.Requests); i++ {
		delay := c.Requests[i].Time.Sub(c.Requests[i-1].Time)
		totalDelay += delay
		count++
	}

	averageDelay := totalDelay / time.Duration(count)
	return averageDelay
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
