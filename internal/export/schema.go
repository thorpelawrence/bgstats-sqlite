package export

import (
	"errors"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Time.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	tt, err := time.Parse(time.DateTime, string(data))
	if err != nil {
		return err
	}
	*t = Time{tt}
	return nil
}

type Model struct {
	Locations []Location `json:"locations"`
	Players   []Player   `json:"players"`
}

type Location struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Modified Time   `json:"modificationDate"`
}

type Player struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Modified Time   `json:"modificationDate"`
}
