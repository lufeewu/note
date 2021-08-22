/*
 * @lc app=leetcode.cn id=676 lang=golang
 *
 * [676] 实现一个魔法字典
 */

// @lc code=start
type MagicDictionary struct {
	dict []string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.dict = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for i := 0; i < len(this.dict); i++ {
		if calc(this.dict[i], searchWord) {
			return true
		}
	}
	return false
}

func calc(dst, src string) bool {
	if len(dst) != len(src) {
		return false
	}
	isFirst := true
	for i := 0; i < len(dst); i++ {
		if src[i] == dst[i] {
			continue
		} else {
			if isFirst {
				isFirst = false
				continue
			}
			return false
		}
	}
	return !isFirst
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end

