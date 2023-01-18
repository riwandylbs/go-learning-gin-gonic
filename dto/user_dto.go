package dto

type UserDTO struct {
	Id      uint64  `json:"id"`
	Name    string  `json:"name"`
	Email   *string `json:"email"`
	Address string  `json:"address"`
	Token   string  `json:"token"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
