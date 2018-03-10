# leetcode 57 题解

#### 题目：

Given a set of *non-overlapping* intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

**Example 1:**
Given intervals `[1,3],[6,9]`, insert and merge `[2,5]` in as `[1,5],[6,9]`.

**Example 2:**
Given `[1,2],[3,5],[6,7],[8,10],[12,16]`, insert and merge `[4,9]` in as `[1,2],[3,10],[12,16]`.

This is because the new interval `[4,9]` overlaps with `[3,5],[6,7],[8,10]`.



参考翻译：

给定一组不重叠的间隔，在间隔中插入新的间隔（如果需要的话就合并）。

您可以认为这些时间间隔最初是根据其开始时间进行排序的。

例1：

给定区间[1,3]，[6,9]，插入并合并[2,5] 结果为 [1,5]，[6,9]。

例2：

给定[1,2]，[3,5]，[6,7]，[8,10]，[12,16]，插入并合并[4,9] 结果为 [1,2]，[3，10]，[12,16]。

`note: 这是因为新的区间[4,9]与[3,5]，[6,7]，[8,10]重叠，因此直接合并为 [3, 10]` 



#### 题解：

总共三种情况，一次遍历即可完成

![img](http://p5d4sh44q.bkt.clouddn.com/2018-03-10%2015-44-00%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

注：newInterval为例一中的[2, 5]，interval为数组遍历的每一个元素

+ 第一种情况

  interval.End < newInterval.Start

  将interval插入result中即可

+ 第二种情况

  interval.Start > newInterval.End

  插入newInterval，并将newInterval赋值为interval

+ 第三种情况

  两个区间有重合部分，只需将连个区间重合即可



代码实现：

```go
// Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
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

// 三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
```

