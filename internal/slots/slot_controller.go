package slots

import (
	"encoding/json"
	"net/http"
)

type SlotController struct {
	ss *SlotService
}

func NewSlotController(ss *SlotService) *SlotController {
	return &SlotController{ss: ss}
}

func (sc *SlotController) Play(w http.ResponseWriter, r *http.Request) {
	res := sc.ss.Play()
	json.NewEncoder(w).Encode(res)
}
