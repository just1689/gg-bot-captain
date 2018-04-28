package model

type Point struct {
	X        float32 `json:"x" bson:"x"`
	Y        float32 `json:"y" bson:"y"`
	Z        float32 `json:"z" bson:"z"`
	Rotation float32 `json:"rotation" bson:"rotation"`
	Radius   float32 `json:"radius" bson:"radius"`
}
