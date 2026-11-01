package problem

import "fmt"

// THIS DIDN'T WORK OUT

type test interface {
	New(start, end int) error
}

type interval struct {
	Interval [2]int
}

func (i *interval) New(start int, end int) error {
	if start > end {
		return fmt.Errorf("start must be less than or equal to end")
	}

	i.Interval[0] = start
	i.Interval[1] = end
	return nil
}

func identifyOverlap(a interval, b interval) bool {
	if a.Interval[0] >= b.Interval[0] && a.Interval[0] <= b.Interval[1] {
		return true
	}

	return true
}
