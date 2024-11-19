package dto

type UserRegisterParams struct {
	Name     string `json:"name" validate:"required,min=2,max=12"`
	Email    string `json:"email" validate:"required,max=255,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}

type UserRegisterResp struct {
	Token         string `json:"token"`
	EmailBeenUsed bool   `json:"emailBeenUsed"`
}

type UserLoginParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResp struct {
	Token                string `json:"token"`
	EmailOrPasswordWrong bool   `json:"emailOrPasswordWrong"`
}
