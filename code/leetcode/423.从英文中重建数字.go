/*
 * @lc app=leetcode.cn id=423 lang=golang
 *
 * [423] 从英文中重建数字
 */

// @lc code=start
func originalDigits(s string) string {
	//arr := []string{
	//	"zero","one","two", "three", "four", "five", "six", "seven", "eight","nine",
	//}
	// u只有four有 (4)
	// x只有six有 (6)
	// g只有eight有 (8)
	// w只有two有 (2)
	// z只有zero有(0)

	// 去除four,two,zero后 o只有one有 (1)
	// 去除four,zero后r只有three有  (3)
	// 去除six后s只有seven有 (7)

	// 剩下的five和nine中 只有five有v  可以分辨5和9

	// 26个字母
	arr := make([]int, 26)
	for _, v := range s {
		arr[v-'a']++
	}

	fmt.Println(arr)
	four := arr['u'-'a']
	shorten(arr, "four", four)
	six := arr['x'-'a']
	shorten(arr, "six", six)
	eight := arr['g'-'a']
	shorten(arr, "eight", eight)
	two := arr['w'-'a']
	shorten(arr, "two", two)
	zero := arr['z'-'a']
	shorten(arr, "zero", zero)

	one := arr['o'-'a']
	shorten(arr, "one", one)
	three := arr['r'-'a']
	shorten(arr, "three", three)
	seven := arr['s'-'a']
	shorten(arr, "seven", seven)

	five := arr['v'-'a']
	shorten(arr, "five", five)

	nine := arr['i'-'a']
	shorten(arr, "nine", nine)

	// fmt.Println(arr, zero, one, two, three, four, five, six, seven, eight, nine)

	res := strings.Repeat("0", zero) + strings.Repeat("1", one) + strings.Repeat("2", two) + strings.Repeat("3", three) + strings.Repeat("4", four) +
		strings.Repeat("5", five) + strings.Repeat("6", six) + strings.Repeat("7", seven) + strings.Repeat("8", eight) + strings.Repeat("9", nine)

	return res
}

// 减少内容
func shorten(arr []int, s string, time int) {
	// fmt.Println(arr, s, time)
	for _, v := range s {
		arr[v-'a'] = arr[v-'a'] - time
	}
}

// @lc code=end

