package client

import (
	"encoding/json"
	"time"
)

type Time struct {
	time.Time `json:",inline"`
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Format(time.RFC3339))
}

func (t Time) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.String()
}

func (t *Time) UnmarshalJSON(p []byte) error {
	if len(p) == 0 {
		return nil
	}

	if err := json.Unmarshal(p, &t.Time); err == nil {
		return nil
	}

	var s string
	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}

	tt, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}
