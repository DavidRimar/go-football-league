package models

type Team struct {
	ID              string `json:"id" bson:"_id"`
	Name            string `json:"name" bson:"name"`
	Stadium         string `json:"stadium" bson:"stadium"`
	StadiumCapacity int64  `json:"stadiumCapacity" bson:"stadiumCapacity"`
}
