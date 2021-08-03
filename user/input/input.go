package input

type LoginInput struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateInput struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Mobile   string `json:"mobile" form:"mobile" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	IDCard   string `json:"id_card" form:"id_card" validate:"required"`
}
