package _1

type StockSpanner struct {
	stack [][2]int // 存储index和val信息
	next  int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: [][2]int{{10000000, -1}},
	}
}

func (this *StockSpanner) Next(price int) int {
	//使用单调栈
	for {
		if price < this.stack[len(this.stack)-1][0] {
			this.stack = append(this.stack, [2]int{price, this.next})
			this.next++
			return this.next - this.stack[len(this.stack)-2][1] - 1
		} else {
			this.stack = this.stack[:len(this.stack)-1]
		}
	}
}
