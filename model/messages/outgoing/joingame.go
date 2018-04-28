package outgoing

//MessageJoinGame is sent to the server to join a game
type MessageJoinGame struct {
	Conversation string `json:"conversation" bson:"conversation"`
	ID           string `json:"id" bson:"id"`
}

//BuildMessageJoinGame creates an object to send to the server
func BuildMessageJoinGame(i string) MessageJoinGame {
	item := MessageJoinGame{
		Conversation: "P_REQUEST_JOIN_GAME",
		ID:           i}
	return item
}
