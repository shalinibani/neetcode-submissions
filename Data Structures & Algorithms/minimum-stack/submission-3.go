type MinStack struct {
 mainStack []int
 minStack []int
 lastMinVal int
}

func Constructor() MinStack {
	return MinStack{lastMinVal: math.MaxInt}
}

func (this *MinStack) Push(val int) {
	this.mainStack = append(this.mainStack, val)
	if this.lastMinVal >= val{
		this.minStack = append(this.minStack, val)
		this.lastMinVal = val
	}
}

func (this *MinStack) Pop() {
 ele := this.mainStack[len(this.mainStack)-1]
 this.mainStack = this.mainStack[:len(this.mainStack)-1]

 if len(this.minStack) > 0{
	if ele == this.minStack[len(this.minStack)-1]{
		this.minStack = this.minStack[:len(this.minStack)-1]
		
		if len(this.minStack) > 0{
			this.lastMinVal = this.minStack[len(this.minStack)-1]
		}else{
			this.lastMinVal = math.MaxInt
		}
	}
 }
}

func (this *MinStack) Top() int {
 return this.mainStack[len(this.mainStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
