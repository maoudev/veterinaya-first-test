package models

type User struct {
	Id        string `json:"id" `
	Rut       string `json:"rut" binding:"required"`
	UserName  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Image_URL string `json:"image_url"`
	Roles     []int  `json:"roles"`
}

type User_SignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
