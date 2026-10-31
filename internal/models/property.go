package models

type Availability string

const (
	Active   Availability = "Available"
	Inactive Availability = "Unavailable"
)

type Rentage string

const (
	Weekly   Rentage = "Weekly"
	Monthly  Rentage = "Monthly"
	Annually Rentage = "Annually"
)

type Property struct {
	ID           uint         `json:"id"`
	Type         string       `json:"type" binding:"required"`
	Location     string       `json:"location" binding:"required"`
	Price        string       `json:"price" binding:"required"`
	RentagePlan  Rentage      `json:"rentage" binding:"required"`
	Description  string       `json:"description" binding:"required"`
	Image        []string     `json:"image" binding:""`
	Reviews      []string     `json:"reviews" `
	Availability Availability `json:"availability" binding:"required"`
	Amenity      Amenities    `json:"amenity" binding:"required"`
}

type Amenities struct {
	Water          bool `json:"water"`
	Parking        bool `json:"parking"`
	Furnished      bool `json:"furnished"`
	Fenced         bool `json:"fence"`
	Prepaid        bool `json:"prepaid"`
	PrivateKitchen bool `json:"private_kitchen"`
}
