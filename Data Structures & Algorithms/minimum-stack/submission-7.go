type MinStack struct {
	stack []data
}

type data struct {
	min int
	value int
}

func Constructor() MinStack {
	return MinStack{stack: []data{}}
}

func (this *MinStack) Push(val int) {
	var d data
	if len(this.stack) != 0 {
		if this.GetMin() > val {
			d.min = val
		} else {
			d.min = this.GetMin()
		}
	} else {
		d.min = val
	}
	d.value = val
	this.stack = append(this.stack, d)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].value
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
