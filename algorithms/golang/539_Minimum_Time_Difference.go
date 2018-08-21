package leetcode

import (
	"sort"
	"strconv"
)

func FindMinDifference(timePoints []string) int {
	sort.Strings(timePoints)

	sHour, sMinute, _ := getHourMinute(timePoints[0])
	eHour, eMinute, _ := getHourMinute(timePoints[len(timePoints)-1])
	result := (24+sHour-eHour)*60 + (sMinute - eMinute)

	for i := 1; i < len(timePoints); i++ {
		eHour, eMinute, _ := getHourMinute(timePoints[i])

		diff := (eHour-sHour)*60 + (eMinute - sMinute)
		if diff < result {
			result = diff
		}
		sHour = eHour
		sMinute = eMinute
	}

	return result
}

func getHourMinute(time string) (int, int, error) {
	hour, err := strconv.Atoi(time[0:2])
	if err != nil {
		return 0, 0, err
	}
	minute, err := strconv.Atoi(time[3:])
	if err != nil {
		return 0, 0, err
	}

	return hour, minute, nil
}
