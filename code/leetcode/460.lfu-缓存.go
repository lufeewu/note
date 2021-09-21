/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
type LFUCache struct {
	freqCache map[int]*list.List
	valCache  map[int]*list.Element
	minFreq   int
	limit     int
}

type Node struct {
	freq int
	key  int
	val  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		freqCache: make(map[int]*list.List, 0),
		valCache:  make(map[int]*list.Element, capacity),
		minFreq:   0,
		limit:     capacity,
	}
}

func (this *LFUCache) Get(key int) int {

	if node, ext := this.valCache[key]; ext {
		data := node.Value.(*Node)
		oldList := this.freqCache[data.freq]
		oldList.Remove(node)
		if data.freq == this.minFreq && oldList.Len() == 0 {
			this.minFreq += 1
		}
		data.freq += 1
		newList, ok := this.freqCache[data.freq]
		if !ok {
			newList = list.New()
			this.freqCache[data.freq] = newList
		}
		e := newList.PushFront(data)
		this.valCache[key] = e
		return data.val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {

	if this.limit <= 0 {
		return
	}
	if node, ext := this.valCache[key]; ext {
		data := node.Value.(*Node)
		oldList := this.freqCache[data.freq]
		oldList.Remove(node)
		if data.freq == this.minFreq && oldList.Len() == 0 {
			this.minFreq += 1
		}
		data.freq += 1
		data.key = key
		data.val = value
		newList, ok := this.freqCache[data.freq]
		if !ok {
			newList = list.New()
			this.freqCache[data.freq] = newList
		}
		e := newList.PushFront(data)
		this.valCache[key] = e
		return
	}
	if len(this.valCache) == this.limit {
		delList, ok := this.freqCache[this.minFreq]
		if ok {
			delNode := delList.Back()
			if delNode != nil {
				dataToDel := delNode.Value.(*Node)
				delList.Remove(delNode)
				delete(this.valCache, dataToDel.key)
				if dataToDel.freq == this.minFreq && delList.Len() == 0 {
					this.minFreq += 1
				}
			}

		}
	}
	newlist, ok := this.freqCache[1]
	if !ok {
		newlist = list.New()
		this.freqCache[1] = newlist
	}
	e := newlist.PushFront(&Node{
		freq: 1,
		key:  key,
		val:  value,
	})
	this.minFreq = 1
	this.valCache[key] = e
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

