/*
 * @lc app=leetcode.cn id=539 lang=golang
 *
 * [539] 最小时间差
 */

// @lc code=start

func findMinDifference(timePoints []string) int {
	var minutes = make([]int, len(timePoints))

	for index, value := range timePoints {
		minutes[index] = timeToMinute(value)
	}
	sort.Ints(minutes)
	var minMin = minutes[0] + 24*60 - minutes[len(minutes)-1]
	for i := 0; i < len(minutes)-1; i++ {
		if minutes[i+1]-minutes[i] < minMin {
			minMin = minutes[i+1] - minutes[i]
		}
	}
	return minMin
}

func timeToMinute(time string) int {
	var times = strings.Split(time, ":")
	var minute int
	if temp, err := strconv.Atoi(times[0]); err == nil {
		minute = minute + temp*60
	}
	if temp, err := strconv.Atoi(times[1]); err == nil {
		minute = minute + temp
	}
	return minute
}

func findMinDifference2(timePoints []string) int {
	sort.Sort(sortTimePoint(timePoints))
	var res, diff = 24 * 60, 0
	for i := 1; i < len(timePoints); i++ {
		diff = timeDiff(timePoints[i], timePoints[i-1])
		if res > diff {
			res = diff
		}
	}
	diff = 24*60 + timeDiff(timePoints[0], timePoints[len(timePoints)-1])
	if res > diff {
		res = diff
	}

	return res
}

func timeDiff(timeA, timeB string) int {
	return timeToMinute(timeA) - timeToMinute(timeB)
}

type sortTimePoint []string

func (s sortTimePoint) Len() int { return len(s) }
func (s sortTimePoint) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortTimePoint) Less(i, j int) bool {
	if timeDiff(s[i], s[j]) <= 0 {
		return true
	}
	return false
}

// @lc code=end

