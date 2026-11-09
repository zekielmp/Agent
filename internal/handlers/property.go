package handlers

import (
	"net/http"
	"strconv"

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

func (h *Handlers) UpdateProperty(w http.ResponseWriter, r *http.Request) {

	req := &dto.Property{}
	if err := resp.ReadJSON(w, r, req); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
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

	if err = services.UpdateProperty(uint(id), prop); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = resp.WriteJSON(w, http.StatusOK, map[string]any{
		"status":  "Success",
		"message": "Property Updated successfully",
	})

}

func (h *Handlers) AddPropertyAmenity(w http.ResponseWriter, r *http.Request) {
	req := &dto.Amenities{}
	if err := resp.ReadJSON(w, r, req); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	amenity := models.Amenities{
		Water:          req.Water,
		Parking:        req.Parking,
		Furnished:      req.Furnished,
		Fenced:         req.Fenced,
		Prepaid:        req.Prepaid,
		PrivateKitchen: req.PrivateKitchen,
	}

	if err = services.AddPropertyAmenity(uint(id), amenity); err != nil {
		_ = resp.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = resp.WriteJSON(w, http.StatusOK, map[string]any{
		"status": "Success",
	})
}
