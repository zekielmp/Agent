package handlers

import (
	"net/http"

	"github.com/zekielmp/Agent/internal/dto"
	resp "github.com/zekielmp/Agent/internal/helpers"
	"github.com/zekielmp/Agent/internal/models"
	"github.com/zekielmp/Agent/internal/services"
)

type Handlers struct{}

func (h *Handlers) SignUp(w http.ResponseWriter, r *http.Request) {
	var dto dto.RegisterRequest

	if err := resp.ReadJSON(w, r, dto); err != nil {
		_ = resp.WriteJSON(w, http.StatusBadRequest, err)
	}

	var user models.User
	res, err := services.CreateUser(&user)
	if err != nil {
		_ = resp.WriteJSON(w, http.StatusInternalServerError, "Failed to sign-up")
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status": "Success",
		"data":   res,
	})

}
