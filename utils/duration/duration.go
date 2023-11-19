package duration

import (
	"encoding/json"
	"time"
)

type Duration struct {
	time.Duration
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler для корректного размаршаливания time.Duration.
func (d *Duration) UnmarshalJSON(data []byte) error {
	var durationString string
	if err := json.Unmarshal(data, &durationString); err != nil {
		return err
	}

	parsedDuration, err := time.ParseDuration(durationString)
	if err != nil {
		return err
	}

	d.Duration = parsedDuration
	return nil
}
