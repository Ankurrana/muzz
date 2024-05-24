package models

type SwipeApiInput struct {
	UserId     int    `json:"user_id" valid:"required"`
	Preference string `json:"preference" valid:"required"`
}

type SwipeApiResponse struct {
	Matched bool `json:"matched"`
	MatchId int  `json:"matchId"`
}
