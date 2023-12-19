package calls

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Request struct {
	UserAgent  string    `json:"userAgent"`
	Data       string    `json:"data"`
	Time       time.Time `json:"time"`
	HTTPMethod string    `json:"method"`
}

func (r Request) Marshal() []byte {
	js, _ := json.Marshal(r)
	return js
}

func (r Request) DumpDataToFile(callName string, outPath string) error {
	f, err := os.Create(fmt.Sprintf("%s/data-%s-%s-%s.json", outPath, callName, r.HTTPMethod, r.Time))
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(r.Data))
	return err
}

func (r Request) DumpToFile(callName string, outPath string) error {
	f, err := os.Create(fmt.Sprintf("%s/%s-%s-%s.json", outPath, callName, r.HTTPMethod, r.Time))
	if err != nil {
		return err
	}

	js, _ := json.MarshalIndent(r, "", "\t")

	_, err = f.Write(js)
	return err
}

func (r Request) DumpToUglyFile(callName string, outPath string) error {
	f, err := os.Create(fmt.Sprintf("%s/ugly-%s-%s-%s.json", outPath, callName, r.HTTPMethod, r.Time))
	if err != nil {
		return err
	}

	js := r.Marshal()

	_, err = f.Write(js)
	return err
}

type Call struct {
	Name       string    `json:"name"`
	Index      int       `json:"-"`
	ID         string    `json:"-"`
	Unanswered bool      `json:"-"`
	Requests   []Request `json:"requests"`
	LastPhone  time.Time `json:"-"`
}

func (c Call) Marshal() []byte {
	js, _ := json.Marshal(c)
	return js
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
