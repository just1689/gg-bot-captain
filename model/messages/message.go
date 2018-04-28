package messages

//ConversationListOfGames is the name of a conversation
const ConversationListOfGames = "S_LIST_OF_GAMES"

//ConversationShareTag is the tag name
const ConversationShareTag = "S_SHARE_TAG"

type Message struct {
	Conversation string `json:"conversation"`
}
