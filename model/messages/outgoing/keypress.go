package outgoing

//KeyUpMessage describes a message to the server
type KeyUpMessage struct {
	Conversation string `json:"conversation" bson:"conversation"`
	Key          string `json:"key" bson:"key"`
}

//KeyDownMessage describes a message to the server
type KeyDownMessage struct {
	Conversation string `json:"conversation" bson:"conversation"`
	Key          string `json:"key" bson:"key"`
}

//BuildKeyUpMessage builds a message object
func BuildKeyUpMessage(key string) KeyUpMessage {
	return KeyUpMessage{Conversation: "P_KU", Key: key}
}

//BuildKeyDownMessage builds a message object
func BuildKeyDownMessage(key string) KeyDownMessage {
	return KeyDownMessage{Conversation: "P_KD", Key: key}
}
