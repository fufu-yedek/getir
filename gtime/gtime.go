package gtime

import (
	"fmt"
	"strings"
	"time"
)

const ResponseFormat = "2006-01-02T15:04:05.999Z07:00"
const RequestFormat = "2006-01-02"

//swagger:strfmt string
type JSONTime time.Time

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		*t = JSONTime{}
		return nil
	}

	tt, err := time.Parse(RequestFormat, s)
	if err != nil {
		return err
	}

	*t = JSONTime(tt)
	return nil
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(ResponseFormat))
	return []byte(stamp), nil
}

func (t JSONTime) ToTime() time.Time {
	return time.Time(t)
}
