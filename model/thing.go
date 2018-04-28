package model

//TableNameThing is the name of the table
const TableNameThing = "Thing"

type Thing struct {
	Point Point  `json:"point" bson:"point"`
	Tag   uint   `json:"tag" bson:"tag"`
	Speed string `json:"speed" bson:"speed"`
}
