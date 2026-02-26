package main

type SlidingWindow struct {
	data []Sample
	size int
}

func NewSlidingWindow(size int) *SlidingWindow {
	return &SlidingWindow{
		data: make([]Sample, 0, size),
		size: size,
	}
}

func (w *SlidingWindow) Add(sample Sample) {
	if len(w.data) >= w.size {
		w.data = w.data[1:]
	}
	w.data = append(w.data, sample)
}

func (w *SlidingWindow) Samples() []Sample {
	return w.data
}