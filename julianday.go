package julianday

import (
	"time"
)

const (
	secondsInADay       = 86400
	unixEpochJulianDate = 2440587.5
)

// JulianDayNumber (JDN) returns the integer of Julian Day on the noon (12:00PM)
// of this date.
func JulianDayNumber(Y int, M time.Month, D int) int64 {
	// https://en.wikipedia.org/wiki/Julian_day#Julian_day_number_calculation
	return (1461*(int64(Y)+4800+(int64(M)-14)/12))/4 +
		(367*(int64(M)-2-12*((int64(M)-14)/12)))/12 -
		(3*((int64(Y)+4900+(int64(M)-14)/12)/100))/4 +
		int64(D) - 32075
}

// JulianDate returns the julian date of time
func JulianDate(t time.Time) float64 {
	return float64(t.Unix())/secondsInADay + unixEpochJulianDate
}
