package ai

import (
	"github.com/just1689/gg-bot-captain/controller"
	"github.com/sirupsen/logrus"
)

func shoot() {
	controller.PressNumber1()
	logrus.Infoln("Pew!")
}
