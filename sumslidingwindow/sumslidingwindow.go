package sumslidingwindow

func NewSumSlidingWindow(size uint) *SumSlidingWindow {
	if size == 0 {
		size = 1
	}

	return &SumSlidingWindow{
		window: make([]int, size),
		size:   size,
	}
}

type SumSlidingWindow struct {
	window []int
	cur    uint
	size   uint
}

func (sw *SumSlidingWindow) Value() int {
	var sum int
	for _, v := range sw.window {
		sum += v
	}
	return sum
}

func (sw *SumSlidingWindow) Add(i int) {
	sw.window[sw.cur] = i
	sw.cur++
	if sw.cur == sw.size {
		sw.cur = 0
	}
}
