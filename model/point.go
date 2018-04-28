package model

type Point struct {
	X        string  `json:"x" bson:"x"`
	Y        string  `json:"y" bson:"y"`
	Z        string  `json:"z" bson:"z"`
	Rotation float32 `json:"rotation" bson:"rotation"`
	Radius   float32 `json:"radius" bson:"radius"`
}
