package model

type Thing struct {
	Point Point   `json:"point" bson:"point"`
	Tag   uint    `json:"tag" bson:"tag"`
	Speed float32 `json:"speed" bson:"speed"`
}
