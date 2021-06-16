package model

type Ratings struct {
	RatingId int64     `json: "id"`
	Rating   []float32 `json:"rating"`
}
