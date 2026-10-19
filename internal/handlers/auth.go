package handlers

import (
	"net/http"

	"github.com/zekielmp/Agent/internal/dto"
	resp "github.com/zekielmp/Agent/internal/helpers"
	"github.com/zekielmp/Agent/internal/models"
	"github.com/zekielmp/Agent/internal/services"
	"github.com/zekielmp/Agent/internal/utils"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) SignUp(w http.ResponseWriter, r *http.Request) {
	var req = &dto.RegisterRequest{}

	if err := resp.ReadJSON(w, r, req); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	var user models.User
	/*Hash password*/
	hashedPassword, _ := utils.Hash(req.Password)
	user.Password = hashedPassword

	/*Create user*/
	res, err := services.CreateUser(&user)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status": "Success",
		"data":   res,
	})

}
