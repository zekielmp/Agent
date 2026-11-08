package handlers

import (
	"net/http"

	"github.com/zekielmp/Agent/internal/dto"
	resp "github.com/zekielmp/Agent/internal/helpers"
	"github.com/zekielmp/Agent/internal/models"
	"github.com/zekielmp/Agent/internal/services"
)

func (h *Handlers) AddProperty(w http.ResponseWriter, r *http.Request) {
	var req = &dto.Property{}
	if err := resp.ReadJSON(w, r, req); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	prop := models.Property{
		Type:         req.Type,
		Location:     req.Location,
		Price:        req.Price,
		RentagePlan:  models.Rentage(req.RentagePlan),
		Description:  string(req.RentagePlan),
		Availability: models.Availability(req.Availability),
	}

	if err := services.AddProperty(prop); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = resp.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":  "Success",
		"message": "Property added successfully",
	})
}
