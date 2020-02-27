package utils

import "time"

const layout = "2006-01-02"

//ConvertStringToTime ...
func ConvertStringToTime(snapshot string) (time.Time, error) {
	t, err := time.Parse(layout, snapshot)
	return t, err
}
