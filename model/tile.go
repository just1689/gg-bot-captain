package model

//TableNameTile is the name of the table
const TableNameTile = "Tile"

//Tile is a skinned point with a model
type Tile struct {
	ID    string `json:"id" bson:"id"`
	Model string `json:"model" bson:"model"`
	Point Point  `json:"point" bson:"point"`
	Skin  string `json:"skin" bson:"skin"`
}
