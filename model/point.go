package model

import "strconv"

//Point is a place that has a rotation and radius
type Point struct {
	X        string  `json:"x" bson:"x"`
	Y        string  `json:"y" bson:"y"`
	Z        string  `json:"z" bson:"z"`
	Rotation float32 `json:"rotation" bson:"rotation"`
	Radius   float32 `json:"radius" bson:"radius"`
}

//SpacialPoint describes a point and rotation
type SpacialPoint struct {
	X        float64
	Y        float64
	Z        float64
	Rotation float32
	Radius   float32
}

//PointToSpacialPoint converts a point to spacial point
func (p Point) PointToSpacialPoint() SpacialPoint {
	item := SpacialPoint{
		Rotation: p.Rotation,
		Radius:   p.Radius}
	item.X, _ = strconv.ParseFloat(p.X, 32)
	item.Y, _ = strconv.ParseFloat(p.Y, 32)
	item.Z, _ = strconv.ParseFloat(p.Z, 32)
	return item

}
