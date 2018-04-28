package incoming

import "github.com/just1689/gg-bot-captain/model"

type MessageShareDynamicThings struct {
	Conversation string        `json:"conversation" bson:"conversation"`
	Things       []model.Thing `json:"things" bson:"things"`
}
