/*
 * @lc app=leetcode.cn id=535 lang=golang
 *
 * [535] TinyURL 的加密与解密
 */

// @lc code=start
type Codec struct {
	numsMap map[string]string
	target  int
}

func Constructor() Codec {
	return Codec{
		numsMap: make(map[string]string),
		target:  1,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	shortUrl := "http://tinyurl.com" + strconv.Itoa(this.target)
	this.numsMap[shortUrl] = longUrl
	this.target++
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.numsMap[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
// @lc code=end

