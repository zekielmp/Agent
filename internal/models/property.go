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
	ID           uint            `json:"id"`
	Type         string          `json:"type" binding:"required"`
	Location     string          `json:"location" binding:"required"`
	Price        string          `json:"price" binding:"required"`
	RentagePlan  Rentage         `json:"rentage" binding:"required"`
	Description  string          `json:"description" binding:"required"`
	Images        []PropertyImage `json:"image" binding:""`
	Reviews      []string        `json:"reviews" `
	Availability Availability    `json:"availability" binding:"required"`
	Amenity      Amenities       `json:"amenity" binding:"required"`
}

type Amenities struct {
	Water          bool `json:"water" default:"false"`
	Parking        bool `json:"parking" default:"false"`
	Furnished      bool `json:"furnished" default:"false"`
	Fenced         bool `json:"fence" default:"false"`
	Prepaid        bool `json:"prepaid" default:"false"`
	PrivateKitchen bool `json:"private_kitchen" default:"false"`
}

// PropertyImage represents an image associated with a product.
type PropertyImage struct {
	ID        uint   `json:"id" `
	ProductID uint   `json:"product_id"`
	URL       string `json:"url" `
	AltText   string `json:"alt_text"`
	IsPrimary bool   `json:"is_primary"`
}
