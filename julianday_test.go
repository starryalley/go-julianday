package julianday

import (
	"math"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestJulianDay(t *testing.T) {
	tables := []struct {
		y int
		m time.Month
		d int
		j int64
	}{
		{2000, time.January, 1, 2451545},
		{2020, time.January, 21, 2458870},
		{1984, time.February, 25, 2445756},
		{1, time.August, 7, 1721644},
		{1993, time.January, 1, 2448989},
		{1995, time.October, 9, 2450000},
		{2023, time.February, 24, 2460000},
	}

	for _, table := range tables {
		j := Number(table.y, table.m, table.d)
		if table.j != j {
			t.Errorf("Julian Day of %d/%d/%d is wrong: %d, expected: %d\n",
				table.y, table.m, table.d, j, table.j)
		}
	}
}

func TestJulianDate(t *testing.T) {
	tables := []struct {
		timeString string
		j          float64
	}{
		{"2000-01-01T00:00:00.0001Z", 2451544.5},
		{"2020-01-22T00:00:00.0001Z", 2458870.5},
		{"1984-02-25T04:30:00.0001Z", 2445755.6875},
		{"1984-03-13T19:50:20.0001Z", 2445773.32662},
		{"0050-07-26T14:42:21.0001Z", 1739529.11274},
	}
	const tolerance = .00001
	opt := cmp.Comparer(func(x, y float64) bool {
		diff := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return (diff / mean) < tolerance
	})

	for _, table := range tables {
		parsed, _ := time.Parse(time.RFC3339, table.timeString)
		j := Date(parsed)
		if !cmp.Equal(j, table.j, opt) {
			t.Errorf("Julian Date of %s is wrong: %.2f expected: %.2f\n",
				table.timeString, j, table.j)
		}
	}
}
