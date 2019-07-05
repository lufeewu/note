 # 简介
 index/suffixarray 标准库通过使用内存中的后缀树实现了对数级时间消耗的子字符串搜索.

 ## 源码
 总计代码 3700 多行，除去测试代码 3100 行
 + type Index
    - func New(data []byte) *Index 生成一个后缀数组
    - func (x *Index) Bytes() []byte
    - func (x *Index) Read(r io.Reader) error
    - func (x *Index) Write(w io.Writer) error
    - func (x *Index) Lookup(s []byte, n int) (result []int)
    - func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)

核心内部函数( go 1.12.5)
+ qsusort(data []byte) []int
+ sortedByFirstByte(data []byte) []int

核心内部函数( go 1.13 用 SA-IS 算法替换了原来的 qsusort )
+ [Two Efficient Algorithms for Linear Time Suffix Array Construction](https://ieeexplore.ieee.org/document/5582081)
    - https://zork.net/~st/jottings/sais.html
+ sais_8_32(text []byte, textMax int, sa, tmp []int32)
+ sais_8_64(text []byte, textMax int, sa, tmp []int64)

+ 后缀数组
    - 通过对字符串的所有后缀进行排序后得到的字符串数组

