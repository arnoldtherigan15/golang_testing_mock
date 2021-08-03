package input

type CarDetailInput struct {
	ID int `param:"id" validate:"required"`
}

// type CarUpdateInput struct {
// 	ID int `param:"id" validate:"required"`
// 	CarType     string    `json:"car_type" form:"car_type" validate:"required"`
// 	Label       string    `json:"label" form:"label" validate:"required"`
// 	Transmision string    `json:"transmision" form:"transmision" validate:"required"`
// 	Fuel        string    `json:"fuel" form:"fuel" validate:"required"`
// 	Stock       int       `json:"stock" form:"stock" validate:"required"`
// 	PricePerDay int       `json:"price_per_day" form:"price_per_day" validate:"required"`
// 	IsActive    bool      `json:"is_active" form:"is_active" validate:"required"``
// }
