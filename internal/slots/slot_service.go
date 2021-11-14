package slots

import (
	"fmt"
	"math/rand"
	"time"
)

type SlotService struct{}

func (ss *SlotService) Play() {
	s := NewSlot()
	rand.Seed(time.Now().UnixNano())
	for i, _ := range s.Reels {
		s.Reels[i] = rand.Intn(10)
	}
	fmt.Println(s.Reels)
}
