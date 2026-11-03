package services

import (
	"context"
	"fmt"
	"time"

	"github.com/zekielmp/Agent/internal/dto"
	"github.com/zekielmp/Agent/internal/models"
)

func AddProperty(p models.Property) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into properties
	 (type,location, price, rentage, description, availability)
	 values ($1, $2, $3, $4, $5, $6)
	 returning id
	`

	err := db.QueryRowContext(ctx, query, p.Type, p.Location, p.Price, p.RentagePlan, p.Description, p.Availability, p.Amenity).Scan(p.ID)
	if err != nil {
		return fmt.Errorf("failed to Add property:%w", err)
	}

	return nil

}

func UpdateProperty(propID uint, p models.Property) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE properties SET 
	type = COALESCE($1, type),
	location = COALESCE($2, location),
	price = COALESCE($3, price),
	rentage = COALESCE($4, rentage),
	description = COALESCE($5, description),
	availability = COALESCE($6, availability)
	`

	_, err := db.ExecContext(ctx, query, p.Type, p.Location, p.Price, p.RentagePlan, p.Description, p.Availability)
	if err != nil {
		return fmt.Errorf("failed to update property:%w", err)
	}
	return nil
}

func AddPropertyAmenity(propID uint, a models.Amenities) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into amenities
	 (property_id, water, parking, furnished, fenced, prepaid, private_kitchen)
	 values ($1, $2, $3, $4, $5, $6, $7)
	 returning id
	`

	err := db.QueryRowContext(ctx, query, propID, a.Water, a.Parking, a.Furnished, a.Fenced, a.Prepaid, a.PrivateKitchen).Scan(a)
	if err != nil {
		return fmt.Errorf("failed to Add property amenity:%w", err)
	}

	return nil

}
func GetProperty(propID uint) (*models.Property, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT * FROM properties WHERE id = $1`

	rows := db.QueryRowContext(ctx, query, propID)
	var p models.Property
	if err := rows.Scan(&p.ID, &p.Type, &p.Location, &p.Price, &p.RentagePlan, &p.Description, &p.Images, &p.Reviews, &p.Availability, &p.Amenity); err != nil {
		return nil, err
	}
	response := models.Property{
		ID:          p.ID,
		Type:        p.Type,
		Location:    p.Location,
		Price:       p.Price,
		RentagePlan: p.RentagePlan,
		Description: p.Description,
		// Image: ,
	}

}

func generatePropertyResponse(p models.Property) *dto.Property {
	image := make([]dto.ImageResponse, len(p.Images))

	for i := range p.Images {
		image[i] = dto.ImageResponse{
			ID:        p.Images[i].ID,
			URL:       p.Images[i].URL,
			AltText:   p.Images[i].AltText,
			IsPrimary: p.Images[i].IsPrimary,
		}
	}

	return &dto.Property{
		ID:           p.ID,
		Type:         p.Type,
		Location:     p.Location,
		Price:        p.Price,
		RentagePlan:  dto.Rentage(p.RentagePlan),
		Description:  p.Description,
		Images:       image,
		Reviews:      p.Reviews,
		Availability: dto.Availability(p.Availability),
		Amenity:      dto.Amenities(p.Amenity),
	}
}
