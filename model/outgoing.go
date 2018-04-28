package model

//MessageRequestJoinServer tells the server about the client
type MessageRequestJoinServer struct {
	Conversation string `json:"conversation"`
	Name         string `json:"name" bson:"name"`
}

//BuildMessageRequestJoinServer creates an object to send to the server
func BuildMessageRequestJoinServer(myName string) MessageRequestJoinServer {
	item := MessageRequestJoinServer{
		Conversation: "P_REQUEST_JOIN_SERVER",
		Name:         myName}
	return item
}
