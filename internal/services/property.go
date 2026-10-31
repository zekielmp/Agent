package services

import (
	"context"
	"fmt"
	"time"

	"github.com/zekielmp/Agent/internal/models"
)

func AddProperty(p models.Property) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into properties
	 (type,location, price, rentage, description, availability, amenity)
	 values ($1, $2, $3, $4, $5, $6, $7)
	 returning id
	`

	err := db.QueryRowContext(ctx, query, p.Type, p.Location, p.Price, p.RentagePlan, p.Description, p.Availability, p.Amenity).Scan(p.ID)
	if err != nil {
		return fmt.Errorf("failed to Add property:%w", err)
	}

	return nil

}
