package request

type Auth struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Level    string
}
