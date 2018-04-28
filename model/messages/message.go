package messages

//ConversationListOfGames is the name of a conversation
const ConversationListOfGames = "S_LIST_OF_GAMES"

//ConversationShareTag is the tag name
const ConversationShareTag = "S_SHARE_TAG"

//ConversationShareDynamicThings is the tag name
const ConversationShareDynamicThings = "S_SHARE_DYNAMIC_THINGS"

//ConversationShareMap is a message from server
const ConversationShareMap = "S_SHARE_MAP"

var ConversationsToIgnore = [...]string{
	"S_SHARE_HEALTH",
	"S_CHANGE_VIEW",
	"S_SCOREBOARD",
	"S_P_LEVEL",
	"S_PLAY_SOUND",
	"S_SHARE_BULLETS",
	"S_SHARE_SPRAY",
	"S_ORB_N",
	"S_PLAYER_LEFT"}

func InConversationsToIgnore(s string) bool {
	for _, b := range ConversationsToIgnore {
		if b == s {
			return true
		}
	}
	return false
}
