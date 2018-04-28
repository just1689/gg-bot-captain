package controller

import "github.com/just1689/gg-bot-captain/model"

var (
	MyTag = 0
)

const ConversationShareTag = "S_SHARE_TAG"

func handleMyTag(tag model.Tag) {
	MyTag = tag.Tag
}
