package models

type Client struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Phone    int    `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
}
