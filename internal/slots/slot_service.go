package slots

import (
	"math/rand"
	"time"
)

type SlotService struct{}

func NewSlotService() *SlotService {
	return &SlotService{}
}

func (ss *SlotService) Play() []int {
	s := NewSlot()
	rand.Seed(time.Now().UnixNano())
	for i, _ := range s.Reels {
		s.Reels[i] = rand.Intn(10)
	}
	return s.Reels
}
