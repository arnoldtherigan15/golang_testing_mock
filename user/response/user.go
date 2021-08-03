package response

// awal nya mo taro di formatter tapi error cyclic import, jadi di pisah folder

import "section9/domain"

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	IDCard   string `json:"id_card"`
	Role     string `json:"role"`
}

func FormatUserResponse(user *domain.User) *UserResponse {
	formatted := UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Mobile:   user.Mobile,
		Address:  user.Address,
		Email:    user.Email,
		IDCard:   user.IDCard,
		Role:     user.Role,
	}
	return &formatted
}

// []users
func FormatUsersResponse(users []*domain.User) []*UserResponse {
	if len(users) == 0 {
		return []*UserResponse{}
	}
	var usersResponse []*UserResponse

	for _, user := range users {
		formatter := FormatUserResponse(user)
		usersResponse = append(usersResponse, formatter)
	}

	return usersResponse
}
