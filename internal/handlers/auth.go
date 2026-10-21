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

	/*Hash password*/
	hashedPassword, _ := utils.Hash(req.Password)
	user := models.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		Password:  hashedPassword,
		LastName:  req.LastName,
		Phone:     req.Phone,
		IsActive:  true,
	}
	/*Create user*/
	err := services.CreateUser(&user)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := dto.AuthResponse{
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			Role:      string(models.UserRoleRenter),
			IsActive:  true,
		},
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status": "Success",
		"data":   res,
	})

}

func (h *Handlers) SignIn(w http.ResponseWriter, r *http.Request) {
	var req = &dto.LoginRequest{}

	if err := resp.ReadJSON(w, r, req); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// var user models.User
	data, err := services.GetUser(*req)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	/*Compare password*/
	if err = utils.Compare(data.Password, req.Password); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := dto.AuthResponse{
		User: dto.UserResponse{
			ID:        data.ID,
			Email:     data.Email,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Phone:     data.Phone,
			Role:      string(models.UserRoleRenter),
			IsActive:  true,
		},
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":  "Success",
		"message": "User signed in successfully",
		"data":    res,
	})

}
