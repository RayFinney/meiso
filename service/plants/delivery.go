package plants

import (
	"context"
	"meiso/models"
)

type Delivery struct {
	plantsService Service
}

func NewDelivery(plantsService Service) Delivery {
	return Delivery{
		plantsService: plantsService,
	}
}

func (d *Delivery) Store(ctx context.Context, plant *models.Plant) error {

}
