/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type IntervalSlice []Interval

func (this IntervalSlice) Len() int {
	return len(this)
}
func (this IntervalSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this IntervalSlice) Less(i, j int) bool {
	return this[i].End < this[j].End
}

func eraseOverlapIntervals(intervals []Interval) int {
    if len(intervals) <= 1 {
        return 0
    }
    
	var newIntervals IntervalSlice
	newIntervals = intervals
	if !sort.IsSorted(newIntervals) {
		sort.Sort(newIntervals)
	}
    
    pre, ans := 0, 1
    for i := 0; i < len(newIntervals); i++ {
        if newIntervals[i].Start < newIntervals[pre].End {
            continue
        }
        pre = i
        ans++
    }
    return len(newIntervals) - ans
}
