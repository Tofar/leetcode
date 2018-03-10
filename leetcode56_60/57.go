package leetcode56_60

// Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	result := []Interval{}
	temp := newInterval
	status := false
	for _, interval := range intervals {
		if temp.Start <= interval.Start {
			if temp.End < interval.Start {
				if !status {
					result = append(result, temp)
					status = true
				}
				result = append(result, interval)

			} else {
				if temp.End < interval.End {
					temp.End = interval.End
				}
			}
		} else {
			if temp.Start <= interval.End {
				if temp.End >= interval.End {
					temp.Start = interval.Start
				} else {
					status = true
					result = append(result, interval)
				}
			} else {
				result = append(result, interval)
			}
		}

	}

	if !status {
		result = append(result, temp)
	}
	return result
}

// 解法2
func insert2(intervals []Interval, newInterval Interval) []Interval {
	result := []Interval{}
	temp := newInterval
	for _, interval := range intervals {
		if interval.End < temp.Start {
			result = append(result, interval)
		} else {
			if interval.Start > temp.End {
				result = append(result, temp)
				temp = interval
			} else {
				temp.Start = If(temp.Start > interval.Start, interval.Start, temp.Start).(int)
				temp.End = If(temp.End < interval.End, interval.End, temp.End).(int)
			}
		}
	}

	return append(result, temp)
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
