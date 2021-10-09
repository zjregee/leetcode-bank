/*
 * @lc app=leetcode.cn id=352 lang=golang
 *
 * [352] 将数据流变为多个不相交区间
 */

// @lc code=start
type SummaryRanges struct {
	*redblacktree.Tree
}


func Constructor() SummaryRanges {
	return SummaryRanges{redblacktree.NewWithIntComparator()}
}


func (ranges *SummaryRanges) AddNum(val int)  {
	interval0, has0 := ranges.Floor(val)
	if has0 && val <= interval0.Value.(int) {
		return
	}
	
	interval1 := ranges.Iterator()
	if has0 {
		interval1 = ranges.IteratorAt(interval0)
	}
	has1 := interval1.Next()
	
	leftAside := has0 && interval0.Value.(int)+1 == val
    rightAside := has1 && interval1.Key().(int)-1 == val
    if leftAside && rightAside {
        interval0.Value = interval1.Value().(int)
        ranges.Remove(interval1.Key())
    } else if leftAside {
        interval0.Value = val
    } else if rightAside {
        right := interval1.Value().(int)
        ranges.Remove(interval1.Key())
        ranges.Put(val, right)
    } else {
        ranges.Put(val, val)
    }
}


func (ranges *SummaryRanges) GetIntervals() [][]int {
	ans := make([][]int, 0, ranges.Size())
	for it := ranges.Iterator(); it.Next(); {
		ans = append(ans, []int{it.Key().(int), it.Value().(int)})
	}
	return ans
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
// @lc code=end

