/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
type LFUCache struct {
	frequencyCache map[int]*list.List
	valueCache     map[int]*list.Element
	minFrequency   int
	capacity       int
}

type Entry struct {
	frequency int
	key       int
	value     int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		frequencyCache: make(map[int]*list.List, 0),
		valueCache:     make(map[int]*list.Element, capacity),
		minFrequency:   0,
		capacity:       capacity,
	}
}
func (this *LFUCache) get(key int) *Entry {
	if elem, ok := this.valueCache[key]; ok {
		data := elem.Value.(*Entry)
		oldList := this.frequencyCache[data.frequency]
		oldList.Remove(elem)
		if data.frequency == this.minFrequency && oldList.Len() == 0 {
			this.minFrequency += 1
		}
		data.frequency += 1
		newList, ok := this.frequencyCache[data.frequency]
		if !ok {
			newList = list.New()
			this.frequencyCache[data.frequency] = newList
		}
		e := newList.PushFront(data)
		this.valueCache[key] = e
		return data
	}
	return nil
}

func (this *LFUCache) Get(key int) int {
	if entry := this.get(key); entry != nil {
		return entry.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}
	if entry := this.get(key); entry != nil {
		entry.value = value
		return
	}

	if len(this.valueCache) == this.capacity {
		delList, ok := this.frequencyCache[this.minFrequency]
		if ok {
			delEle := delList.Back()
			if delEle != nil {
				dataToDel := delEle.Value.(*Entry)
				delList.Remove(delEle)
				delete(this.valueCache, dataToDel.key)
				if dataToDel.frequency == this.minFrequency && delList.Len() == 0 {
					this.minFrequency += 1
				}
			}

		}
	}

	newlist, ok := this.frequencyCache[1]
	if !ok {
		newlist = list.New()
		this.frequencyCache[1] = newlist
	}
	e := newlist.PushFront(&Entry{
		frequency: 1,
		key:       key,
		value:     value,
	})
	this.minFrequency = 1
	this.valueCache[key] = e
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

