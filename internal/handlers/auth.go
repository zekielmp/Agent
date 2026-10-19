package handlers

import (
	"net/http"

	"github.com/zekielmp/Agent/internal/dto"
	resp "github.com/zekielmp/Agent/internal/helpers"
)

type Handlers struct{}

func (h *Handlers) SignUp(w http.ResponseWriter, r *http.Request) {
	var dto dto.RegisterRequest

	if err := resp.ReadJSON(w, r, dto); err != nil {
		_ = resp.WriteJSON(w, http.StatusBadRequest, err)
	}
	

}
