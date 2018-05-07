package outgoing

//KeyUpMessage describes a message to the server
type KeyUpMessage struct {
	Conversation string
	Key          string
}

//KeyDownMessage describes a message to the server
type KeyDownMessage struct {
	Conversation string
	Key          string
}

//BuildKeyUpMessage builds a message object
func BuildKeyUpMessage(key string) KeyUpMessage {
	return KeyUpMessage{Conversation: "P_K_UP", Key: key}
}

//BuildKeyDownMessage builds a message object
func BuildKeyDownMessage(key string) KeyDownMessage {
	return KeyDownMessage{Conversation: "P_K_DOWN", Key: key}
}
