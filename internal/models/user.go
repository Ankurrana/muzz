package models

type UserApiInput struct {
	Name     string  `json:"name" valid:"required"`
	Email    string  `json:"email" valid:"email"`
	Password string  `json:"password"`
	Age      int     `json:"age" valid:"range(1|99)"`
	Gender   string  `json:"gender" valid:"matches(male|female|lgbtqp)"`
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

type DiscoveredUser struct {
	User
	Id             int
	DistanceFromMe float64 `json:"distance_from_me"`
	Score          float64 `json:"attractiveness"`
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
