package mp4

import (
	"fmt"
	"math"
)

// GetDurationString Helper function to print a duration value in the form H:MM:SS.MS
func GetDurationString(duration uint32, timescale uint32) string {
	// TODO: zero padding

	durationSec := float64(duration) / float64(timescale)

	hours := math.Floor(durationSec / 3600)
	durationSec -= hours * 3600

	minutes := math.Floor(durationSec / 60)
	durationSec -= minutes * 60

	msec := durationSec * 1000
	durationSec = math.Floor(durationSec)
	msec -= durationSec * 1000
	msec = math.Floor(msec)

	tmpl := fmt.Sprintf("%.0f:%.0f:%.0f:%.0f", hours*10, minutes*10, durationSec, msec)
	return tmpl
}
