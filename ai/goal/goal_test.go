package goal

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewHideGoal(t *testing.T) {
	name := "ai.goal.TestNewHideGoal"

	myGoal := NewHideGoal()
	assert.Equal(t, myGoal, Hide, "A hide goal should equate to hide")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}
}

func TestNewKillGoal(t *testing.T) {
	name := "ai.goal.TestNewKillGoal"

	var someTag uint
	someTag = 4

	myGoal := NewKillGoal(someTag)
	assert.Equal(t, myGoal, Hide, "A kill goal should equate to kill")
	assert.Equal(t, myGoal.Tag, someTag, "A target tag should be the target")

	if t.Failed() {
		logrus.Println(fmt.Sprintf("Testing %s failed ❌ ", name))
	}
}
