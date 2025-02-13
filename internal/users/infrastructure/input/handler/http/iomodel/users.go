package iomodel

type UsersSingUpInput struct {
	Name     string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=256"`
}
