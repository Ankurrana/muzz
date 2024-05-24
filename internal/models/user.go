package models

type UserApiInput struct {
	Name     string  `json:"name" valid:"required"`
	Email    string  `json:"email" valid:"email"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
	Gender   string  `json:"gender"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

type User struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
}

type DBUser struct {
	User
	Id       int
	Password string
}

type LoginRequest struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"required"`
}
