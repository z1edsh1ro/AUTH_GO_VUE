package dto

type RegisterRequestDTO struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty" binding:"required,min=2"`
	Email    string `json:"email,omitempty" bson:"email,omitempty" binding:"required,email,min=6"`
	Password string `json:"password,omitempty" bson:"password,omitempty" binding:"required,min=6"`
}
