package DS

import "math"

type MinStack struct {
	mainStack, minStack []int
}

func Constructor() MinStack {
	return MinStack{
		mainStack: []int{},
		minStack:  []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.mainStack = append(this.mainStack, val)
	v := this.minStack[len(this.minStack)-1]
	if val < v {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, v)
	}
}

func (this *MinStack) Pop() {
	this.mainStack = this.mainStack[:len(this.mainStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.mainStack[len(this.mainStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
