package slots

type Slot struct {
	Reels []int
}

func NewSlot() *Slot {
	s := &Slot{Reels: make([]int, 3)}
	return s
}
