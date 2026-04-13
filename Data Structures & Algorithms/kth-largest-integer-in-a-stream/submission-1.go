type KthLargest struct {
    k int
    nums []int
}


func Constructor(k int, nums []int) KthLargest {
    kl := KthLargest{
        k: k,
        nums: []int{},
    }
    for _, val := range nums {
        kl.Add(val)
    }
    return kl
}

func (this *KthLargest) Swap(i, j int) {
	this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

func (this *KthLargest) Less(i, j int) bool {
	return this.nums[i] < this.nums[j]
}

func (this *KthLargest) Len() int {
	return len(this.nums)
}

func (this *KthLargest) Pop() int {
	res := this.nums[0]
	last := this.nums[this.Len()-1]
	this.nums = this.nums[:this.Len()-1]

	if this.Len() == 0 {
		return res
	}

	this.nums[0] = last

	i := 0
	for 2*i+1 < this.Len() {
		smallest := i
		left := 2*i + 1
		right := 2*i + 2
		if left < this.Len() && this.Less(left, smallest) {
			smallest = left
		}
		if right < this.Len() && this.Less(right, smallest) {
			smallest = right
		}
		if smallest == i {
			break
		}
		this.Swap(i, smallest)
		i = smallest
	}
	return res
}

func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	i := this.Len() - 1
	// percolate up
	for i > 0 && this.Less(i, (i-1)/2) {
		this.Swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
	if this.Len() > this.k {
		this.Pop()
	}
	return this.nums[0]
}
