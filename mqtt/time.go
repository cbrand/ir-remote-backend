package mqtt

import (
	"encoding/json"
	"time"

	"github.com/btubbs/datetime"
)

// Time includes a time and allows to format and unmarshal json
type Time struct {
	time.Time
}

// ToIso8601 converts the time into an isoformatted string
func (t Time) ToIso8601() string {
	return t.Time.Format(time.RFC3339)
}

// UnmarshalJSON deserializes the JSON byte stream to a datetime object
func (t Time) UnmarshalJSON(data []byte) error {
	var jsonString string
	err := json.Unmarshal(data, &jsonString)
	if err != nil {
		return err
	}
	parsedTime, err := datetime.Parse(jsonString, time.UTC)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
