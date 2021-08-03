package input

type Order struct {
	EstimatedDays int  `json:"estimated_days" validate:"required"`
	WithDriver    bool `json:"with_driver" validate:"required"`
	CarID         int  `json:"car_id" validate:"required"`
	UserID        int  `json:"user_id" validate:"required"`
}

type DoneInput struct {
	OrderID int `json:"order_id" validate:"required"`
}
