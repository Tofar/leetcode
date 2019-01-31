package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func (this Intervals) Len() int           { return len(this) }
func (this Intervals) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this Intervals) Less(i, j int) bool { return this[i].Start < this[j].Start }

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(Intervals(intervals))
	res := []Interval{}
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if cur.End < intervals[i].Start {
			res = append(res, cur)
			cur = intervals[i]
		} else if cur.End < intervals[i].End {
			cur.End = intervals[i].End
		}
	}
	res = append(res, cur)
	return res
}

func main() {
	fmt.Println()
}
