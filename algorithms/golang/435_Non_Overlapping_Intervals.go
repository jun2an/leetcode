package leetcode
/**
 * Definition for an interval.
 * type Interval struct {
 *       Start int
 *       End   int
 * }
 */

/*
 * 贪心策略：按照intervals.End字段sort即可，无需管理intervals.Start字段。
 */

func eraseOverlapIntervals(intervals []Interval) int {
    if len(intervals) == 0 {
        return 0
    }
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i].End < intervals[j].End {
            return true
        }
        return false
    })

    result := 0
    start := intervals[0]
    for i := 1; i < len(intervals); i++ {
        if intervals[i].Start < start.End {
            result++
        } else {
            start = intervals[i]
        }
    }

    return result
}
